package model

import (
	"github.com/bongcheon/go-blog/db/mongodb"
)

type Article struct {
	mongodb.Document `bson:",inline"`
	Subject string
	Body string
}

