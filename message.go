package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
	"github.com/gomodels/dbutil"
)

type Message struct {
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title  string `json:"title"`
	Text string `json:"text"`
	CreatedIn   time.Time `json:"createdIn"`
	CreatedBy bson.ObjectId `json:"createdBy"`
	Room  bson.ObjectId `json:"room"`
}

type Messages []Message


// Creates a new user
func NewMessage() (*Message) {
	return &Message{}
}

func (message *Message) Persist(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	message.Id = bson.NewObjectId()
	err = c.Insert(message)
	if err != nil {
		return err
	}
	return nil
}

func (messages Messages) FindAll(c *mgo.Collection,id bson.ObjectId) (Messages, error) {
	defer dbutil.CloseSession(c)
	err := c.Find(bson.M{"room": id}).All(&messages)
	if err != nil {
		return messages, err
	}
	return messages, nil
}