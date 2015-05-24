package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/bongcheon/go-blog/api"
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
	router.Run(":8080")
}
