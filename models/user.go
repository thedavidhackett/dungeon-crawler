package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//User ..
type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email string `bson:"email" json:"email"`
	Name string `bson:"name" json:"name"`
	VerifiedEmail bool `bson:"verifiedEmail" json:"-"`
}
