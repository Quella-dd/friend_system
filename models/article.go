package models

import "time"

type ArticleManager struct {}

type Article struct {
	Content string
	UserID string
	CreateTime time.Time
}

func (m *ArticleManager) GetUserArticles(id string) ([]Article, error) {
	return nil, nil
}

func (m *ArticleManager) CreateArticle(article Article) (*Article, error) {
	return nil, nil
}

func (m *ArticleManager) GetArticle(id string) (*Article, error) {
	return nil, nil
}

func (m *ArticleManager) DeleteArticle(id string) error {
	return nil
}


func NewArticleManager() *ArticleManager{
	return &ArticleManager{}
}