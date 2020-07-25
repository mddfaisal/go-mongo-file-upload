package email

import (
	"encoding/json"
	"fmt"
	"services/db"
	"services/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// Email email struct
type Email struct {
	ID   string `json:"_id"`
	Muid string `json:"muid"`
	HTML string `json:"html"`
}

// New new email
func (e *Email) Create() string {
	dbh := db.GetDb()
	result, err := dbh.Collection.InsertOne(*dbh.Ctx, e)
	if err != nil {
		fmt.Println(utils.Trace())
		panic(err)
	}
	e.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return e.ID
}

// FindOne find one document
func (e *Email) FindOne() ([]byte, error) {
	dbh := db.GetDb()
	objID, _ := primitive.ObjectIDFromHex(e.ID)
	err := dbh.Collection.FindOne(*dbh.Ctx, bson.M{"_id": objID}).Decode(&e)
	str, _ := json.Marshal(e)
	return str, err
}
