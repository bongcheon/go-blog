package api

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"time"
	"net/http"
	"github.com/bongcheon/go-blog/model"
	"github.com/bongcheon/go-blog/db/mongodb"
	"gopkg.in/mgo.v2/bson"
)

func GetArticles(c *gin.Context) {
	var err error
	var page, pagesize int
	page, err = strconv.Atoi(c.DefaultQuery("page", "1"))
	pagesize, err = strconv.Atoi(c.DefaultQuery("pagesize", "8"))

	if page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if pagesize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	var skipcnt = (page - 1) * pagesize

	var results []bson.M

	err = mongodb.GetCollection("Article").Find(bson.M{}).Sort("-createdat").Skip(skipcnt).Limit(pagesize).All(&results)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, results)
}

func GetArticle(c *gin.Context) {
	id := c.Params.ByName("id")

	article := &model.Article{}

	err := mongodb.GetCollection("Article").FindByStrId(id, article)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "NotFound",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": article.GetId(),
		"subject": article.Subject,
		"body": article.Body,
		"type": article.Type,
		"createdAt": article.CreatedAt,
	})
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
		return
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

func UpdateArticle(c *gin.Context) {
	id := c.Params.ByName("id")

	var update ArticleJSON
	c.Bind(&update)

	err := mongodb.GetCollection("Article").UpdateByStrId(id, bson.M{"$set":update})
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":id,
			"subject":update.Subject,
			"body":update.Body,
		})
	}
}

