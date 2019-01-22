package dao

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	. "db_connect/model"
)

type MyDao struct {
	Server string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "test_collection"
)

func (m *MyDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//insert user
func (m *MyDao) InsertUser(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

func (m *MyDao) GetAllUsers() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}