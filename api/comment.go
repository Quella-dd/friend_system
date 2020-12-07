package api

import (
	"friend_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteComment(c *gin.Context) {
	userID := c.GetHeader("userID")
	id := c.Param("id")
	if err := models.ManagerEnv.DeleteComment(userID, id); err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}
