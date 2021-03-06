package model

import (
	"github.com/bongcheon/go-blog/db/mongodb"
	"time"
)

type ArticleType string

const (
	ArticleText ArticleType = "Text"
	ArticleLink ArticleType = "Link"
	ArticleVideo ArticleType = "Video"
	ArticleAbout ArticleType = "About"
)

type Article struct {
	mongodb.Document `bson:",inline"`
	Subject string
	Body string
	Type ArticleType
	CreatedAt time.Time
}

func (a *Article) SetType(t ArticleType) {
	switch t {
		case ArticleText, ArticleLink, ArticleVideo, ArticleAbout:
			a.Type = t;
		default:
			a.Type = "";
	}
}

