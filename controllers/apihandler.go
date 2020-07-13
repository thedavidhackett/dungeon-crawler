package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"dungeon-crawler/models"
	"dungeon-crawler/middlewares"
	"dungeon-crawler/websocket"
)

//APIHandler ..
type APIHandler struct {
	db mongo.Database
	store *sessions.CookieStore
}

//NewAPIHandler ..
func NewAPIHandler(db mongo.Database, store *sessions.CookieStore,) *APIHandler {
	return &APIHandler{
		db: db,
		store: store,
	}
}

//GetDungeonsCreated ..
func (h *APIHandler) GetDungeonsCreated(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middlewares.AuthenticatedUserKey).(models.User)

	collection := h.db.Collection("dungeons")
	findOptions := options.Find()

	var dungeons []*models.Dungeon = make([]*models.Dungeon, 0)

	cur, err := collection.Find(context.TODO(), bson.M{"userId": user.ID}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem models.Dungeon
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		dungeons = append(dungeons, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	json.NewEncoder(w).Encode(dungeons)
}

//CreateDungeon ..
func (h *APIHandler) CreateDungeon(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middlewares.AuthenticatedUserKey).(models.User)

	collection := h.db.Collection("dungeons")
	var d models.Dungeon
	err := json.NewDecoder(r.Body).Decode(&d)
	d.Characters = make([]models.Character, 0)
	if (err != nil) {
		//todo error handling
	}
	d.UserID = user.ID
	result, err := collection.InsertOne(context.TODO(), d)
	if (err != nil) {
		//todo error handling
	}
	json.NewEncoder(w).Encode(result)

}

//GetDungeon ..
func (h *APIHandler) GetDungeon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if (err != nil) {
		//todo error handling
	}

	collection := h.db.Collection("dungeons")
	var d models.Dungeon

	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&d)
	if (err != nil) {
		//todo error handling
	}

	json.NewEncoder(w).Encode(d)
}

//EditDungeon ..
func (h *APIHandler) EditDungeon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if (err != nil) {
		//todo error handling
	}

	collection := h.db.Collection("dungeons")
	var d models.Dungeon
	err = json.NewDecoder(r.Body).Decode(&d)

	result, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": objectID}, d)
	if (err != nil) {
		//todo error handling
	}

	json.NewEncoder(w).Encode(result)
}

//GetDungeonCharacters ..
func (h *APIHandler) GetDungeonCharacters(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	objectID, err := primitive.ObjectIDFromHex(id)
	if (err != nil) {
		//todo error handling
	}

	collection := h.db.Collection("characters")
	findOptions := options.Find()

	var characters []*models.Character = make([]*models.Character, 0)

	cur, err := collection.Find(context.TODO(), bson.M{"dungeonId": objectID}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem models.Character
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		characters = append(characters, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	json.NewEncoder(w).Encode(characters)
}


//CreateCharacter ..
// func (h *APIHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
// 	user := r.Context().Value(middlewares.AuthenticatedUserKey).(models.User)

// 	collection := h.db.Collection("characters")
// 	var c models.Character
// 	err := json.NewDecoder(r.Body).Decode(&c)
// 	if (err != nil) {
// 		//todo error handling
// 	}
// 	c.UserID = user.ID
// 	result, err := collection.InsertOne(context.TODO(), c)
// 	if (err != nil) {
// 		//todo error handling
// 	}
// 	if newID, ok := result.InsertedID.(primitive.ObjectID); ok {
// 		c.ID = newID
// 		json.NewEncoder(w).Encode(c)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(map[string]string{"error": "something went wrong"})
// }

//WebSocketEndpoint ..
func (h *APIHandler) WebSocketEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomID := params["id"]

	ws, err := websocket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c := &websocket.Connection{Send: make(chan models.Dungeon), Ws: ws}
	s := websocket.Subscription{Conn: c, Room: roomID, Db: h.db}
	websocket.Hub.Register <- s
	go s.WritePump()
	go s.ReadPump()

}
