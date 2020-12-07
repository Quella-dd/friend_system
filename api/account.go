package api

import (
	"friend_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var body struct {
		Name string
		Password string
	}
	if err := c.ShouldBind(&body); err != nil {
		panic(err)
		return
	}
	if err := models.ManagerEnv.Login(); err != nil {
		panic(err)
		return
	}
}

func Logout(c *gin.Context)  {
}


func ListFriend(c *gin.Context) {
	id := c.GetHeader("userID")
	if users, err := models.ManagerEnv.ListFriend(id); err != nil {
		panic(err)
	}  else {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

func AddFriend(c *gin.Context) {
	id := c.GetHeader("userID")
	friendID := c.Param("id")
	if err := models.ManagerEnv.AddFriend(id, friendID); err != nil {
		panic(err)
	}  else {
		c.JSON(http.StatusOK, nil)
	}
}

func DeleteFriend(c *gin.Context) {
	id := c.GetHeader("userID")
	friendID := c.Param("id")
	if err := models.ManagerEnv.DeleteFriend(id, friendID); err != nil {
		panic(err)
	}  else {
		c.JSON(http.StatusOK, nil)
	}
}

func SearchUsers(c *gin.Context) {
	name := c.Param("name")
	if users, err := models.ManagerEnv.SearchUsers(name); err != nil {
		panic(err)
	}  else {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}