package api

import (
	"friend_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRequests(c *gin.Context) {
	userID := c.GetHeader("userID")
	if requests, err := models.ManagerEnv.ListUserRequest(userID); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"request": requests,
		})
	}
}

func DeleteRequest(c *gin.Context) {
	id := c.Param("id")
	if err := models.ManagerEnv.DeleteRequest(id); err !=  nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}