package models

type CommentManager struct {}

type Comment struct {
	UserID  string
	Content string
}

func (m *CommentManager) GetComments(id string) ([]Comment, error) {
	return nil, nil
}

func (m *CommentManager) DeleteComment(useID string, id string) error {
	return nil
}

func NewCommentManager() *CommentManager {
	return &CommentManager{}
}
