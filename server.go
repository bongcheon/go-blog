package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bongcheon/go-blog/api"
	"github.com/bongcheon/go-blog/config"
)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	router.GET("/articles/:id", api.GetArticle)
	router.PUT("/articles/:id", api.UpdateArticle)
	router.DELETE("/articles/:id", api.DeleteArticle)
	router.POST("/articles", api.PostArticle)
	router.Run(":" + config.Get("server_port"))
}
