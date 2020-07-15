package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Character ..
type Character struct {
	// ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
	Npc bool `bson:"npc" json:"npc"`
	// DungeonID primitive.ObjectID `bson:"dungeonId" json:"dungeonId"`
	UserID primitive.ObjectID `bson:"userId" json:"userId"`
	PositionX int `bson:"positionX" json:"positionX"`
	PositionY int `bson:"positionY" json:"positionY"`
	Image string `bson:"image" json:"image"`
	Initiative int `bson:"initiative" json:"initiative"`
}
