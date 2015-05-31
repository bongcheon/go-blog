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

func (c *Collection) UpdateById(id bson.ObjectId, update interface{}) error {
	collection := c.internalGetCollection()
	err := collection.UpdateId(id, update)
	return err
}

func (c *Collection) UpdateByStrId(id string, update interface{}) error {
	if bson.IsObjectIdHex(id) == false {
		return nil
	}

	err := c.UpdateById(bson.ObjectIdHex(id), update)
	return err
}

func (c *Collection) RemoveByStrId(id string) error {
	if bson.IsObjectIdHex(id) == false {
		return nil
	}

	return c.RemoveById(bson.ObjectIdHex(id))
}

func (c *Collection) RemoveById(id bson.ObjectId) error {
	collection := c.internalGetCollection()
	return collection.RemoveId(id)
}

func (c *Collection) FindOne(query interface{}, doc interface{}) error {
	collection := c.internalGetCollection()
	err := collection.Find(query).One(doc)

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

