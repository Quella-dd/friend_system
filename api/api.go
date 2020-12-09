package api

import (
	"fmt"
	"friend_system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var routers = []Router{
	{Method: http.MethodPost, Path: "/api/account/login", Handler: Login},
	{Method: http.MethodPost, Path: "/api/account/registry", Handler: Registry},

	{Method: http.MethodGet, Path: "/api/friends", Handler: ListFriend},
	{Method: http.MethodPost, Path: "/api/friend/:id", Handler: AddFriend},
	{Method: http.MethodDelete, Path: "/api/friend/:id", Handler: DeleteFriend},
	{Method: http.MethodGet, Path: "/api/friends/:name", Handler: SearchUsers},

	{Method: http.MethodGet, Path: "/api/articles", Handler: ListArticles},
	{Method: http.MethodPost, Path: "/api/article", Handler: CreateArticle},
	{Method: http.MethodGet, Path: "/api/article/:id", Handler: GetArticle},
	{Method: http.MethodDelete, Path: "/api/article/:id", Handler: DeleteArticle},

	{Method: http.MethodPost, Path: "/api/comment/:id", Handler: CreateComment},
	{Method: http.MethodDelete, Path: "/api/comment/:id", Handler: DeleteComment},

	{Method: http.MethodGet, Path: "/api/requests/", Handler: ListRequests},
	{Method: http.MethodPost, Path: "/api/requests/:id", Handler: AckRequest},
	{Method: http.MethodDelete, Path: "/api/request/:id", Handler: DeleteRequest},
}

func InitRouters() {
	engine := gin.Default()
	for _, r := range routers {
		engine.Handle(r.Method, r.Path, r.Handler)
	}
	port := fmt.Sprintf(":%s", models.ManagerConfig.Port)
	engine.Run(port)
}