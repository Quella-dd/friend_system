package models

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ManagerEnv *Manger
var ManagerConfig ManagerIni

type Manger struct {
	*gorm.DB
	*AccountManager
	*ArticleManager
	*CommentManager
	*RequestManager
	*PhotoManager
}

type ManagerIni struct {
	Mysql
	Port string
	FilePath string
	SecretKey string
}

type Mysql struct {
	UserName string
	Password string
	DBName string
	Host string
}

func InitManage() {
	ManagerEnv =  &Manger{
		AccountManager: NewAccountManager(),
		ArticleManager: NewArticleManager(),
		CommentManager: NewCommentManager(),
		RequestManager: NewRequestManager(),
		PhotoManager: NewPhotoManager(),
	}
}

func LoadInit() {
	cfg, err := ini.Load("/home/len/go/src/friend_system/config.ini")
	if err != nil {
		panic(err)
	}
	ManagerConfig.Port = cfg.Section("").Key("Port").String()
	ManagerConfig.FilePath = cfg.Section("").Key("FilePath").String()
	ManagerConfig.SecretKey = cfg.Section("").Key("SecretKey").String()

	ManagerConfig.Mysql.UserName = cfg.Section("mysql").Key("username").String()
	ManagerConfig.Mysql.Password = cfg.Section("mysql").Key("password").String()
	ManagerConfig.Mysql.DBName = cfg.Section("mysql").Key("dbname").String()
	ManagerConfig.Mysql.Host = cfg.Section("mysql").Key("host").String()
}

func initMysql() {
	sql := ManagerConfig.Mysql
	url := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", sql.UserName, sql.Password, sql.Host, sql.DBName)
	if db, err := gorm.Open("mysql", url); err != nil {
		panic(err)
	} else {
		ManagerEnv.DB = db
	}
}

func initDataTable() {
	ManagerEnv.DB.AutoMigrate(&User{})
	ManagerEnv.DB.AutoMigrate(&Article{})
	ManagerEnv.DB.AutoMigrate(&Comment{})
	ManagerEnv.DB.AutoMigrate(&Request{})
	ManagerEnv.DB.AutoMigrate(&PhotoRepository{})
	ManagerEnv.DB.AutoMigrate(&Photo{})
}

// InitFileServe to save avatar and photo
func initFileServe() {
	if err := os.Mkdir(ManagerConfig.FilePath, 0777); err != nil && !os.IsExist(err) {
		panic(err)
	}
}

func init() {
	InitManage()
	LoadInit()
	initMysql()
	initDataTable()
	initFileServe()
}