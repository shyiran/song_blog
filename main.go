package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"song_blog/models"
	_ "song_blog/routers"
)


func init() {

	models.Init()
	web.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	web.Run()
}