package models

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gomodels/dbutil"
)

type Room struct {
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name"`
	Tag string `json:"tag"`
	Pwd   string `json:"pwd"`
	CreatedBy bson.ObjectId `json:"createdBy"`
	Users  []bson.ObjectId `json:"users"`
}

type Rooms []Room

// Creates a new room
func NewRoom() (*Room) {
	return &Room{Name: "", Tag: "", Pwd: ""}
}

func (room *Room) Persist(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	room.Id = bson.NewObjectId()
	err = c.Insert(room)
	log.Println("Sala", room.Tag, "inserida")
	if err != nil {
		return err
	}
	return nil
}

func (room *Room) Merge(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	err = c.Update(bson.M{"_id": room.Id}, &room)
	log.Println("Sala", room.Tag, "atualizada")
	if err != nil {
		return err
	}
	return nil
}

func (room *Room) Remove(c *mgo.Collection) error {
	var err error
	defer dbutil.CloseSession(c)
	err = c.Remove(bson.M{"_id": room.Id})
	log.Println("Sala", room.Tag, "removida")
	if err != nil {
		return err
	}
	return nil
}

func (room *Room) FindById(c *mgo.Collection, id bson.ObjectId) error {
	defer dbutil.CloseSession(c)
	err := c.Find(bson.M{"_id": id}).One(&room)
	if err != nil {
		return err
	}
	return nil
}

func (rooms Rooms) FindByUserId(c *mgo.Collection, id bson.ObjectId) (Rooms, error) {
	defer dbutil.CloseSession(c)
	err := c.Find(bson.M{"users": bson.M{"$in":[]bson.ObjectId{id}}}).All(&rooms)
	if err != nil {
		return rooms,err
	}
	return rooms, nil
}

func (rooms Rooms) FindAll(c *mgo.Collection) (Rooms, error) {
	defer dbutil.CloseSession(c)
	err := c.Find(bson.M{}).All(&rooms)
	if err != nil {
		return rooms, err
	}
	return rooms, nil
}
