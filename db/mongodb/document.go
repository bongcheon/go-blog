package mongodb

import (
	"gopkg.in/mgo.v2/bson"
)

type SuperDocument interface {
	GetId() bson.ObjectId
	SetId(bson.ObjectId)
}

type Document struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"_id"`
}

func (d *Document) GetId() bson.ObjectId {
	return d.Id
}

func (d *Document) SetId(id bson.ObjectId) {
	d.Id = id
}

