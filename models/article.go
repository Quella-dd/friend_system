package models

import (
	"github.com/jinzhu/gorm"
)

type ArticleManager struct {}

func NewArticleManager() *ArticleManager{
	return &ArticleManager{}
}

type Article struct {
	gorm.Model
	Content string `form:"content"`
	UserID string
}

type ArticleInfo struct {
	 Article
	 Comments []Comment
}

func (m *ArticleManager) GetUserArticles(id string) ([]ArticleInfo, error) {
	var articleInfos []ArticleInfo
	var articles []Article
	if err := ManagerEnv.DB.Where("user_id = ?", id).Find(&articles).Error; err != nil {
		return nil, err
	}
	for _, art := range articles {
		var articleInfo ArticleInfo
		articleInfo.Article = art
		if comments, err := ManagerEnv.GetComments(art.ID); err != nil {
			articleInfo.Comments = []Comment{}
		} else {
			articleInfo.Comments = comments
		}
		articleInfos = append(articleInfos, articleInfo)
	}
	return articleInfos, nil
}

func (m *ArticleManager) CreateArticle(article Article) (*Article, error) {
	if err := ManagerEnv.DB.Save(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (m *ArticleManager) GetArticle(id string) (*ArticleInfo, error) {
	var article Article
	if err := ManagerEnv.DB.First(&article, id).Error; err != nil {
		return nil, err
	}
	var articleInfo ArticleInfo
	articleInfo.Article = article

	if comments, err := ManagerEnv.GetComments(article.ID); err != nil {
		articleInfo.Comments = []Comment{}
	} else {
		articleInfo.Comments = comments
	}

	return &articleInfo, nil
}

// 当用户删除说说时，应删除Article底下的所有comments
func (m *ArticleManager) DeleteArticle(id string) error {
	go func() {
		if err := ManagerEnv.DB.Where("article_id = ?", id).Delete(&Comment{}).Error; err != nil {
			// TODO: 是否需要忽略删除错误, 或者给每个用户创建一个定时job,每隔24h清除一次无用的comments
		}
	}()
	return ManagerEnv.DB.Where("id = ?", id).Delete(Article{}).Error
}