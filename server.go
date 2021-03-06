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

	mongoHost := os.Getenv("MONGO_HOST")
	if mongoHost == "" {
		mongoHost = config.Get("mongodb_host")
	}

	mongodb.Init(mongoHost, config.Get("mongodb_db"))

	router := gin.Default()

	// CORS
	router.Use(func(c *gin.Context) {
		// Run this on all requests
		// Should be moved to a proper middleware
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,POST,DELETE")
		c.Next()
	})

	router.OPTIONS("/*cors", func(c *gin.Context) {
		// Empty 200 response
	})
	// CORS END

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "go-blog api server!")
	})
	router.GET("/users/:username", api.GetUser)
	router.POST("/users", api.AddUser)
	router.GET("/articles", api.GetArticles)
	router.GET("/articles/:id", api.GetArticle)
	router.PUT("/articles/:id", api.UpdateArticle)
	router.DELETE("/articles/:id", api.DeleteArticle)
	router.POST("/articles", api.PostArticle)
	router.Run(":" + port)
}
