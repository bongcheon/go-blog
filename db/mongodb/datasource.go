package mongodb

import (
	"gopkg.in/mgo.v2"
)

type Connection struct {
	HostName string
	DatabaseName string
	Session *mgo.Session
}

func (conn *Connection) GetCollection(name string) *Collection {
	return &Collection {
		Name: name,
		Connection: conn,
	}
}

func (conn *Connection) Open() {
	var err error
	conn.Session, err = mgo.Dial(conn.HostName)
	if err != nil {
		panic(err)
	}
}

func (conn *Connection) Close() {
	conn.Session.Close()
}

func GetConnection() (*Connection) {
	conn := &Connection {
		HostName: "localhost", //TODO
		DatabaseName: "go-blog-dev", //TODO
	}
	conn.Open()
	return conn
}

