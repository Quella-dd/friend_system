package models

import "github.com/jinzhu/gorm"

type RequestManager struct {}

func NewRequestManager() *RequestManager {
	return &RequestManager{}
}

type Request struct {
	gorm.Model
	UserID string
	AddID string
	Content string
}

func (m *RequestManager) Createrequest(id, addID, content string) error {
	var request Request
	request.UserID = id
	request.AddID = addID
	request.Content = content
	return ManagerEnv.DB.Save(&request).Error
}

func (m *RequestManager) ListUserRequest(id string) ([]Request, error) {
	var requests []Request
	if err := ManagerEnv.DB.Where("userID = ?", id).Find(&requests).Error; err != nil {
		return nil, err
	}
	return requests, nil
}

func (m *RequestManager) DeleteRequest(id string) error {
	if err := ManagerEnv.DB.Where("id = ?", id).Delete(Request{}).Error; err != nil {
		return err
	}
	return nil
}

func (m *RequestManager) AckRequest(id string) error {
	var request Request
	if err := ManagerEnv.DB.Where("id = ?", id).First(&request).Error; err != nil {
		return err
	}
	return ManagerEnv.AckRequet(request.UserID, request.AddID)
}