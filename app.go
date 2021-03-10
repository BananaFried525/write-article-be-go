package main

import (
	"fmt"
	"os"
	"write-article/src/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(cors.Default())
	router(r)
	r.NoRoute(func(c *gin.Context) {
		ret := make(map[string]interface{})
		ret["error"] = "Page not found"
		c.JSON(404, ret)
	})
	r.Run(getPort())
}

func router(router *gin.Engine) {
	index := router.Group("/")
	{
		index.GET("/", controllers.Index)
	}

	article := router.Group("/article")
	{
		article.POST("/article", controllers.PostCreateArticle)
		article.GET("/list", controllers.GetArticleList)
	}
}

func getPort() string {
	p := os.Getenv("port")
	port := ""
	if p != "" {
		port = p
	} else {
		port = "8080"
	}
	return fmt.Sprintf(`:%v`, port)
}
