package controllers

import (
	"log"
	"write-article/src/db"

	"github.com/gin-gonic/gin"
)

var (
	conn = db.ConnectDB()
	ret  = make(map[string]interface{})
)

func Index(c *gin.Context) {
	ret["data"] = "Hello"
	log.Println(ret)
	c.JSON(200, ret)
	return
}
