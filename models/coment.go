package models

import (
	"github.com/jinzhu/gorm"
)

type CommentManager struct {}

func NewCommentManager() *CommentManager {
	return &CommentManager{}
}

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
	return ManagerEnv.DB.Save(&comment).Error
}

func (m *CommentManager) DeleteComment(id string) error {
	return ManagerEnv.DB.Delete(&Comment{}, id).Error
}

