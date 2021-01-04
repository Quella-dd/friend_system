package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type AccountManager struct {}

func NewAccountManager() *AccountManager{
	return &AccountManager{}
}

type AddUserOptions struct {
	Content string `form:"content"`
}

type User struct {
	gorm.Model
	Name      string `form:"name"`
	PassWord  string `form:"password"`
	Email     string `form:"email"`
	Relations RelationStruct `gorm:"type:json"`
}

type UpdateOptions struct {
	Name string `form:"name"`
	Email string `form:"email"`
}

type RelationStruct []string

func (u RelationStruct) Value() (driver.Value, error) {
	b, err := json.Marshal(u)
	return string(b), err
}

func (u *RelationStruct) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), u)
}

// ------------------- Login and Registry
func (m *AccountManager) Login(user *User) (string, error) {
	if err := ManagerEnv.DB.Find(user).Error; err != nil {
		return "", errors.New("recode has been exist")
	}

	return GenerateToken(user.Name, strconv.Itoa(int(user.ID)), 24 * time.Hour)
}

// 创建用户禁止用户名称重复
func (m *AccountManager) Registry(user *User) error {
	if err := ManagerEnv.DB.Where("name = ?", user.Name).Error; err != nil {
		return ManagerEnv.DB.Save(user).Error
	}
	return errors.New("name already taken")
}


// ------------------------ update and delete
func (m *AccountManager) UpdateAccount(id string, options *UpdateOptions) error {
	if user, err := m.GetUser(id); err  != nil {
		return err
	} else {
		return ManagerEnv.DB.Model(&user).Update(options).Error
	}
}

func (m *AccountManager) DeleteAccount(id string) error {
	return ManagerEnv.DB.Delete("id = ?", id).Delete(&User{}).Error
}


// ----------------------- Friends
func (m *AccountManager) ListFriend(id string) ([]User, error) {
	var users []User
	user, err := m.GetUser(id)
	if err != nil {
		return nil, err
	}
	if err := ManagerEnv.DB.Find(&users, []string(user.Relations)).Error; err != nil {
		return nil, fmt.Errorf("ListFriend error:", err)
	}
	return users, nil
}

func (m *AccountManager) AddFriend(id, addID, content string) error {
	if content == "" {
		if user, err := ManagerEnv.GetUser(addID); err != nil {
			return err
		} else {
			content = fmt.Sprintf("'I'm %s\n", user.Name)
		}
	}
	return ManagerEnv.Createrequest(id, addID, content)
}

func (m *AccountManager) DeleteFriend(id, friendID string) error {
	self, _:= m.GetUser(id)
	if _,  err := m.GetUser(friendID); err != nil {
		return errors.New("user not found")
	}
	for i, v := range self.Relations {
		if v == friendID {
			self.Relations = append(self.Relations[:i], self.Relations[i+1:]...)
		}
	}
	return ManagerEnv.DB.Model(&self).Update("relations", self.Relations).Error
}

func (m *AccountManager) SearchUsers(name string) ([]User, error) {
	var users []User
	query := fmt.Sprintf("%%%s%%", name)
	if err := ManagerEnv.DB.Where("name LIKE ?", query).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (m *AccountManager) GetUser(id string) (*User, error) {
	var user User
	if err := ManagerEnv.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, fmt.Errorf("%s, %+v\n", id, err.Error())
	}
	return &user, nil
}

func (m *AccountManager) AckRequet(id, friendID string) error {
	self, _:= m.GetUser(id)
	if _,  err := m.GetUser(friendID); err != nil {
		return errors.New("user not found")
	}
	self.Relations = append(self.Relations, friendID)

	// 更新user relations and request's status, 因此要开启一个事务，保证数据一致性
	tx := ManagerEnv.DB.Begin()

	if err := tx.Model(&self).Update("relations", self.Relations).Error; err != nil {
		tx.Rollback()
		return err
	}

	request, err := ManagerEnv.GetRequest(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&request).Updates(Request{Status: true}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}