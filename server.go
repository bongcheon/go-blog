package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/bongcheon/go-blog/api"
	"github.com/bongcheon/go-blog/config"
	"github.com/bongcheon/go-blog/db/mongodb"
)

func main() {

	// Use development as default
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Get("server_port")
	}

	mongodb.Init(config.Get("mongodb_host"), config.Get("mongodb_db"))

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	router.GET("/articles/:id", api.GetArticle)
	router.PUT("/articles/:id", api.UpdateArticle)
	router.DELETE("/articles/:id", api.DeleteArticle)
	router.POST("/articles", api.PostArticle)
	router.Run(":" + port)
}
