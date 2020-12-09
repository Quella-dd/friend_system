package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
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

type RelationStruct []string

func (u RelationStruct) Value() (driver.Value, error) {
	b, err := json.Marshal(u)
	return string(b), err
}

func (u *RelationStruct) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), u)
}

func (m *AccountManager) Login(user User) error {
	if err := ManagerEnv.DB.Find(user).Error; err != nil {
		return errors.New("recode has been exist")
	}
	return nil
}

func (m *AccountManager) Registry(user *User) error {
	if err := ManagerEnv.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (m *AccountManager) ListFriend(id string) ([]User, error) {
	var users []User
	user, err := m.GetUser(id)
	if err != nil {
		return nil, err
	}
	if err := ManagerEnv.DB.Find(&users, []string(user.Relations)).Error; err != nil {
		return nil, err
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
		return nil, err
	}
	return &user, nil
}

func (m *AccountManager) AckRequet(id, friendID string) error {
	self, _:= m.GetUser(id)
	if _,  err := m.GetUser(friendID); err != nil {
		return errors.New("user not found")
	}
	self.Relations = append(self.Relations, friendID)
	return ManagerEnv.DB.Model(&self).Update("relations", self.Relations).Error
}