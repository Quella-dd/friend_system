package api

import (
	"friend_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRequests(c *gin.Context) {
	userID := c.GetString("userID")
	if requests, err := models.ManagerEnv.ListUserRequest(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"request": requests,
		})
	}
}

func AckRequest(c *gin.Context) {
	id := c.Param("id")
	if err := models.ManagerEnv.AckRequest(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func DeleteRequest(c *gin.Context) {
	id := c.Param("id")
	if err := models.ManagerEnv.DeleteRequest(id); err !=  nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}