package mongodb

import (
	"gopkg.in/mgo.v2"
)

var (
	db *mgo.Database
)

func Init(host string, dbname string) {
	session, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	db = session.DB(dbname)
}

func GetDB() *mgo.Database {
	return db
}

func GetCollection(cname string) *Collection {
	return &Collection{
		Name: cname,
	}
}
