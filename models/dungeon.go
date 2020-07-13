package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Dungeon ..
type Dungeon struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
	Height string `bson:"height" json:"height"`
	Width string `bson:"width" json:"width"`
	UserID primitive.ObjectID `bson:"userId" json:"userId"`
	Characters []Character `bson:"characters" json:"characters"`
}
