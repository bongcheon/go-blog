package model

import (
	"github.com/bongcheon/go-blog/db/mongodb"
)

type User struct {
	mongodb.Document `bson:",inline"`
	Name string
	Username string
	Twitter string
	Disqus string
	GitHub string
}

