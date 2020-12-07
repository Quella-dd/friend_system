package main

import (
	"friend_system/api"
	_ "friend_system/models"
)

func main() {
	api.InitRouters()
}