package api

import (
	"friend_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context) {
	userID := c.GetString("userID")
	id := c.Param("id")

	var comment models.Comment
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	comment.UserID = userID
	comment.ArticleID = id

	if err := models.ManagerEnv.CreateComment(comment); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	if err := models.ManagerEnv.DeleteComment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
}
