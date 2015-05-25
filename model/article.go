package model

import (
	"github.com/bongcheon/go-blog/db/mongodb"
	"time"
)

type Article struct {
	mongodb.Document `bson:",inline"`
	Subject string
	Body string
	CreatedAt time.Time
}

