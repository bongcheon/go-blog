package mongodb

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Collection struct {
	Name string
}

func (c *Collection) internalGetCollection() *mgo.Collection {
	return GetDB().C(c.Name)
}

func (c *Collection) Save(doc SuperDocument) error {
	collection := c.internalGetCollection()
	_, err := collection.UpsertId(doc.GetId(), doc)

	if err != nil {
		return err
	}

	return nil
}

func (c *Collection) FindByStrId(id string, doc interface{}) error {
	if bson.IsObjectIdHex(id) == false {
		return errors.New("Invalid ObjectId")
	}

  return c.FindById(bson.ObjectIdHex(id), doc)
}

func (c *Collection) FindById(id bson.ObjectId, doc interface{}) error {
	collection := c.internalGetCollection()
	err := collection.FindId(id).One(doc)

	if err != nil {
		return err
	}

	return nil
}

