package api

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"net/http"
	"github.com/bongcheon/go-blog/model"
	"github.com/bongcheon/go-blog/db/mongodb"
	"gopkg.in/mgo.v2/bson"
)

func GetArticle(c *gin.Context) {
	id := c.Params.ByName("id")

	article := &model.Article{}

	err := mongodb.GetCollection("Article").FindByStrId(id, article)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "NotFound",
		})
		return;
	}

	c.JSON(http.StatusOK, gin.H{
		"id": article.GetId(),
		"subject": article.Subject,
		"body": article.Body,
		"type": article.Type,
		"createdAt": article.CreatedAt,
	})
}

func UpdateArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"id":id})//TODO
}

func DeleteArticle(c *gin.Context) {
	id := c.Params.ByName("id")

	err := mongodb.GetCollection("Article").RemoveByStrId(id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{})
}

type ArticleJSON struct {
	Subject string `json:"subject" binding:"required"`
	Body string `json:"body" binding:"required"`
	Type model.ArticleType `json:"type" binding:"required"`
}

func PostArticle(c *gin.Context) {

	var json ArticleJSON
	c.Bind(&json)

	article := &model.Article{
		Subject: json.Subject,
		Body: json.Body,
		CreatedAt: time.Now(),
	}
	article.SetId(bson.NewObjectId())
	article.SetType(json.Type)
	if article.Type == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return;
	}

	err := mongodb.GetCollection("Article").Save(article)
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":article.GetId(),
			"subject":article.Subject,
			"body":article.Body,
			"type":article.Type,
			"createdAt":article.CreatedAt,
		})
	}
}

