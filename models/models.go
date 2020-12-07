package models

import (
	"fmt"
	"gopkg.in/ini.v1"

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
}

type ManagerIni struct {
	 Mysql
	 Port string
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
	}
}

func LoadInit() {
	cfg, err := ini.Load("/home/len/go/src/friend_system/config.ini")
	if err != nil {
		panic(err)
	}
	ManagerConfig.Port = cfg.Section("").Key("Port").String()

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

func init() {
	InitManage()
	LoadInit()
	initMysql()
}