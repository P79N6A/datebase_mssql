package main

import (
	"code.byted.org/gopkg/logs"
	"fmt"
	"github.com/gin-gonic/gin"
	"task/store"
)

func Ping(c *gin.Context) {
	logs.Info("a sample app log")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Select(c *gin.Context) {
	resp, err:= store.GetG()
	if err != nil {
		c.JSON(500, gin.H{ "message":fmt.Sprintln("get member failed, err=%v",err),})
	    return
	}
	c.JSON(200, gin.H{ "message":resp,})
}
