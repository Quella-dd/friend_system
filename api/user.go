package api

import (
	"friend_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	if err := models.ManagerEnv.Login(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "username or password error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"response": "login success",
		})
	}
}

func Registry(c *gin.Context)  {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	if err := models.ManagerEnv.Registry(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func ListFriend(c *gin.Context) {
	id := c.GetHeader("userID")
	if users, err := models.ManagerEnv.ListFriend(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}  else {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

func AddFriend(c *gin.Context) {
	id := c.GetHeader("userID")
	addID := c.Param("id")

	var option models.AddUserOptions
	_ = c.ShouldBind(&option)
	if err := models.ManagerEnv.AddFriend(id, addID, option.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}  else {
		c.JSON(http.StatusOK, nil)
	}
}

func DeleteFriend(c *gin.Context) {
	id := c.GetHeader("userID")
	friendID := c.Param("id")
	if err := models.ManagerEnv.DeleteFriend(id, friendID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
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