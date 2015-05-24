package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"id":id})
}

func UpdateArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"id":id})
}

func DeleteArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"id":id})
}

func PostArticle(c *gin.Context) {
	c.String(http.StatusUnauthorized, "not authorized")
}

