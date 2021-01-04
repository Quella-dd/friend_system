package models

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PhotoManager struct{}

func NewPhotoManager() *PhotoManager {
	return &PhotoManager{}
}

type PhotoRepository struct {
	gorm.Model

	Name        string `form:"name"`
	Description string `form:"description"`
	UserID      string
}

type PhotoRepositoryInfo struct {
	PhotoRepository
	Count int
}

type Photo struct {
	Name         string
	RepositoryID string
}

func (m *PhotoManager) ListRepository(userID string) ([]PhotoRepositoryInfo, error) {
	var repositoriyInfoes []PhotoRepositoryInfo
	var repositories []PhotoRepository
	if err := ManagerEnv.DB.Where("UserID = ?", userID).Find(&repositories).Error; err != nil {
		return nil, err
	}

	for _, repository := range repositories {
		count, _ := GetRepositoryPhotoCount(repository)
		repositoriyInfoes = append(repositoriyInfoes, PhotoRepositoryInfo{
			PhotoRepository: repository,
			Count: count,
		})
	}
	return repositoriyInfoes, nil
}

func (m *PhotoManager) GetRepository(id string) (*PhotoRepository, error) {
	var repository PhotoRepository
	if err := ManagerEnv.DB.Where("id = ?", id).Find(&repository).Limit(1).Error; err != nil {
		return nil, err
	}
	return &repository, nil
}

func (m *PhotoManager) CreateRepository(repository *PhotoRepository, userID string) error {
	repository.UserID = userID
	if err := ManagerEnv.DB.Save(&repository).Error; err != nil {
		return err
	}
	return nil
}

// update Description, Name
func (m *PhotoManager) UpdateRepository(id string, repository *PhotoRepository) error {
	return ManagerEnv.DB.Where("id = ?", id).Update(&repository).Error
}

func (m *PhotoManager) DeleteRepository(id string) error {
	if err := ManagerEnv.DB.Where("id = ?", id).Delete(&PhotoRepository{}).Error; err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s", ManagerConfig.FilePath, id)
	// 当删除相册时，同时删除相册下面的图片
	return DeletePhotoFile(path)
}

func (m *PhotoManager) UploadPhoto(c *gin.Context, repositoryID, photoName string) error {
	dir := fmt.Sprintf("%s/%s", ManagerConfig.FilePath, repositoryID)
	if err := os.Mkdir(dir, 0777); err != nil && !os.IsExist(err) {
		return err
	}

	path := fmt.Sprintf("%s/%s/%s", ManagerConfig.FilePath, repositoryID, photoName)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, c.Request.Body)
	defer c.Request.Body.Close()
	return err
}

func (m *PhotoManager) DeletePhoto(c *gin.Context, repositoryID, photoName string) error {
	path := fmt.Sprintf("%s/%s/%s", ManagerConfig.FilePath, repositoryID, photoName)
	return DeletePhotoFile(path)
}

func DeletePhotoFile(path string) error {
	return os.Remove(path)
}

func GetPath(c *gin.Context) (string, error) {
	userID := c.GetHeader("userID")
	repositoryID := c.Param("repositoryID")
	photoID := c.Param("id")
	if userID == "" || repositoryID == "" || photoID == "" {
		return "", errors.New("param not illegel")
	}
	return fmt.Sprintf("%s/%s/%s/\n", userID, repositoryID, photoID), nil
}

func GetRepositoryPhotoCount(repository PhotoRepository) (int, error) {
	path := fmt.Sprintf("%s/%+v", ManagerConfig.FilePath, repository.ID)

	filepathNames, err := filepath.Glob(filepath.Join(path,"*"))
	if err != nil {
		return 0, err
	}
	return len(filepathNames), nil
}
