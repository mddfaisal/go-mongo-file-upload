package email

import (
	"encoding/json"
	"services/db"

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
func (e *Email) New() (string, error) {
	dbh := db.GetDb()
	result, err := dbh.Collection.InsertOne(*dbh.Ctx, e)
	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

// FindOne find one document
func (e *Email) FindOne() ([]byte, error) {
	dbh := db.GetDb()
	objID, _ := primitive.ObjectIDFromHex(e.ID)
	err := dbh.Collection.FindOne(*dbh.Ctx, bson.M{"_id": objID}).Decode(&e)
	str, _ := json.Marshal(e)
	return str, err
}
