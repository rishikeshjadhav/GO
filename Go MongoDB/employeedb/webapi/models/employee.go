package models

import "gopkg.in/mgo.v2/bson"

// Model for employee
type Employee struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Code       string        `bson:"code" json:"code"`
	Name       string        `bson:"name" json:"name"`
	Department string        `bson:"department" json:"department"`
}
