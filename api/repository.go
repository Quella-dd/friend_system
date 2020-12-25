package api

import (
	"friend_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRepository(c *gin.Context) {
	repository, err := models.ManagerEnv.ListRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"repository": repository})
	}
}

func GetRepository(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	if repository, err := models.ManagerEnv.GetRepository(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"repository": repository})
	}
}

func CreateRepository(c *gin.Context) {
	var repository models.PhotoRepository
	userID := c.GetString("userID")
	if err := c.ShouldBind(&repository); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	if err := models.ManagerEnv.CreateRepository(&repository, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func UpdateRepository(c *gin.Context) {
	var repository models.PhotoRepository
	if err := c.ShouldBind(&repository); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	if err := models.ManagerEnv.UpdateRepository(&repository); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func DeleteRepository(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	if err := models.ManagerEnv.DeleteRepository(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusOK, nil)
	}

}

func UploadPhoto(c *gin.Context) {
	repositoryID := c.Param("id")
	photoName := c.Param("photoName")

	if err := models.ManagerEnv.UploadPhoto(c, repositoryID, photoName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func DeletePhoto(c *gin.Context) {
	repositoryID := c.Param("id")
	photoName := c.Param("photoName")

	if err := models.ManagerEnv.DeletePhoto(c, repositoryID, photoName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	} else {
		c.JSON(http.StatusOK, nil)
	}
}