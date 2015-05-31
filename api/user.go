package api

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/bongcheon/go-blog/model"
	"github.com/bongcheon/go-blog/db/mongodb"
	"gopkg.in/mgo.v2/bson"
)

func GetUser(c *gin.Context) {
	username := c.Params.ByName("username")

	user := &model.User{}

	err := mongodb.GetCollection("User").FindOne(bson.M{"username":username}, user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "NotFound",
		})
		return;
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
		"username": user.Username,
		"twitter": user.Twitter,
		"disqus": user.Disqus,
		"github": user.GitHub,
	})
}

type UserJSON struct {
	Name string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Twitter string `json:"twitter"`
	Disqus string `json:"disqus"`
	GitHub string `json:"github"`
}

func AddUser(c *gin.Context) {

	var json UserJSON
	c.Bind(&json)

	user := &model.User{
		Name: json.Name,
		Username: json.Username,
		Twitter: json.Twitter,
		Disqus: json.Disqus,
		GitHub: json.GitHub,
	}
	user.SetId(bson.NewObjectId())

	err := mongodb.GetCollection("User").Save(user)
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"name": user.Name,
			"username": user.Username,
			"twitter": user.Twitter,
			"disqus": user.Disqus,
			"github": user.GitHub,
		})
	}
}

