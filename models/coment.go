package models

import (
	"github.com/jinzhu/gorm"
)

type CommentManager struct {}

// A 评论 B, Content
type Comment struct {
	gorm.Model
	ArticleID string
	UserID  string
	Content string `form:"content"`
}

func (m *CommentManager) GetComments(id interface{}) ([]Comment, error) {
	var comments []Comment
	if err := ManagerEnv.DB.Where("article_id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (m *CommentManager) CreateComment(comment Comment) error {
	if err := ManagerEnv.DB.Save(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (m *CommentManager) DeleteComment(id string) error {
	if err := ManagerEnv.DB.Delete(&Comment{}, id).Error; err != nil {
		return err
	}
	return nil
}

func NewCommentManager() *CommentManager {
	return &CommentManager{}
}
