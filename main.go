package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"task/store"
)

func main() {
	r := gin.Default()
	logger.SetLogger("log/log.json")
	store.InitDB()
	r.GET("/ping", Ping)
	r.GET("/select", Select)
	r.Run() // listen and serve on 0.0.0.0:8080
}




