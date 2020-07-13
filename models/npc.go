package models

//Npc ..
type Npc struct {
	Name string `bson:"name" json:"name"`
	PositionX int `bson:"positionX" json:"positionX"`
	PositionY int `bson:"positionY" json:"positionY"`
}
