package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type RequestManager struct {}

func NewRequestManager() *RequestManager {
	return &RequestManager{}
}

type Request struct {
	gorm.Model
	UserID string
	AddID string
	Content string
	Status bool
}

func (m *RequestManager) Createrequest(id, addID, content string) error {
	var request Request
	request.UserID = id
	request.AddID = addID
	request.Content = content
	request.Status = false
	return ManagerEnv.DB.Save(&request).Error
}

func (m *RequestManager) GetRequest(id string) (*Request, error) {
	var request Request
	if err := ManagerEnv.DB.Where("id = ?", id).Find(&request).Error; err != nil {
		return nil, err
	}
	return &request, nil
}

func (m *RequestManager) ListUserRequest(id string) ([]Request, error) {
	var requests []Request
	if err := ManagerEnv.DB.Where("add_id = ?", id).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (m *RequestManager) DeleteRequest(id string) error {
	return ManagerEnv.DB.Where("id = ?", id).Delete(Request{}).Error
}

func (m *RequestManager) AckRequest(id, userID string) error {
	var request Request
	if err := ManagerEnv.DB.Where("id = ?", id).First(&request).Error; err != nil {
		return err
	}

	if userID != request.AddID {
		return fmt.Errorf("%s not permission to resolve %s's request", userID, request.AddID)
	}

	return ManagerEnv.AckRequet(request.UserID, request.AddID)
}