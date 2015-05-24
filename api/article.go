package api

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/bongcheon/go-blog/model"
	"github.com/bongcheon/go-blog/db/mongodb"
	"gopkg.in/mgo.v2/bson"
)

func GetArticle(c *gin.Context) {
	id := c.Params.ByName("id")

	article := &model.Article{
	}
	err := mongodb.GetCollection("Article").FindById(bson.ObjectIdHex(id), article)
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"id":article.GetId(),"subject":article.Subject,"body":article.Body})
	}
}

func UpdateArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"id":id})//TODO
}

func DeleteArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{"id":id})//TODO
}

func PostArticle(c *gin.Context) {

	//FIXME
	article := &model.Article{
		Subject:"New subject",
		Body:"New body",
	}
	article.SetId(bson.NewObjectId())

	err := mongodb.GetCollection("Article").Save(article)
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"id":article.GetId(),"subject":article.Subject,"body":article.Body})
	}
}

