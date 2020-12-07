package api

import (
	"friend_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListArticles(c *gin.Context) {
	userID := c.GetHeader("userID")
	if articles, err := models.ManagerEnv.GetUserArticles(userID); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"articles": articles,
		})
	}
}

func CreateArticle(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBind(&article); err != nil {
		panic(err)
		return
	}
	if a, err := models.ManagerEnv.CreateArticle(article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"article": a,
		})
	}
	return
}

func GetArticle(c *gin.Context) {
	id := c.Param("id")
	if article, err := models.ManagerEnv.ArticleManager.GetArticle(id); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"article": article,
		})
	}
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	if err := models.ManagerEnv.ArticleManager.DeleteArticle(id); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}