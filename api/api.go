package api

import (
	"errors"
	"fmt"
	"friend_system/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type Router struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var routers = []Router{
	// 可以内置一个管理员用户，管理员用户对其他用户拥有管理的权限
	{Method: http.MethodPost, Path: "/api/account/login", Handler: Login},
	{Method: http.MethodPost, Path: "/api/account/registry", Handler: Registry},

	{Method: http.MethodPut, Path: "/api/account/:id", Handler: UpdateAccount},
	{Method: http.MethodDelete, Path: "/api/account/:id", Handler: DeleteAccount},

	{Method: http.MethodGet, Path: "/api/friends", Handler: ListFriend},
	{Method: http.MethodPost, Path: "/api/friend/:id", Handler: AddFriend},
	{Method: http.MethodDelete, Path: "/api/friend/:id", Handler: DeleteFriend},
	{Method: http.MethodGet, Path: "/api/friends/:name", Handler: SearchUsers},

	// 说说详细信息，包括用户评论
	{Method: http.MethodGet, Path: "/api/articles", Handler: ListArticles},
	{Method: http.MethodPost, Path: "/api/article", Handler: CreateArticle},
	{Method: http.MethodGet, Path: "/api/article/:id", Handler: GetArticle},
	{Method: http.MethodDelete, Path: "/api/article/:id", Handler: DeleteArticle},

	// 添加评论
	{Method: http.MethodPost, Path: "/api/comment/:id", Handler: CreateComment},
	{Method: http.MethodDelete, Path: "/api/comment/:id", Handler: DeleteComment},

	// 添加用户请求
	{Method: http.MethodGet, Path: "/api/requests/", Handler: ListRequests},
	{Method: http.MethodPost, Path: "/api/request/:id", Handler: AckRequest},
	{Method: http.MethodDelete, Path: "/api/request/:id", Handler: DeleteRequest},

	// 相册，包括创建、删除、更新、详情
	{Method: http.MethodGet, Path: "/api/repositories", Handler: ListRepository},
	{Method: http.MethodGet, Path: "/api/repository/:id", Handler: GetRepository},
	{Method: http.MethodPost, Path: "/api/repository", Handler: CreateRepository},
	{Method: http.MethodPut, Path: "/api/repository/:id", Handler: UpdateRepository},
	{Method: http.MethodDelete, Path: "/api/repository/:id", Handler: DeleteRepository},

	// TODO: 监控每个用户固定的文件的大小，当用户上传图片时，如果总的文件大小大于设定值，不允许用户上传

	// 上传图片，删除图片
	{Method: http.MethodPost, Path: "/api/photo/repository/:id/:photoName", Handler: UploadPhoto},
	{Method: http.MethodDelete, Path: "/api/photo/:photoName/repository/:id", Handler: DeletePhoto},

	// TODO: 留言板， 包括创建、删除、评论、详情
}

func InitRouters() {
	engine := gin.Default()
	for _, r := range routers {
		engine.Handle(r.Method, r.Path, validateToken(r.Handler))
	}
	port := fmt.Sprintf(":%s", models.ManagerConfig.Port)
	log.Fatalln(engine.Run(port))
}

func validateToken(f func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.URL.Path, "/login") && !strings.Contains(c.Request.URL.Path, "/registry") {
			tokenAuth := c.GetHeader("token")
			if tokenAuth == "" {
				c.JSON(http.StatusUnauthorized, gin.H {
					"error": errors.New("http.StatusUnauthorized"),
				})
				return
			} else {
				t, err := jwt.ParseWithClaims(tokenAuth, &models.LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(models.ManagerConfig.SecretKey), nil
				})

				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H {
						"error": errors.New("http.StatusUnauthorized"),
					})
					return
				}
				if claims, ok := t.Claims.(*models.LoginClaims); ok && t.Valid {
					c.Set("userName", claims.UserName)
					c.Set("userID", claims.UserID)
				}
			}
		}
		f(c)
	}
}