package controllers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
    "crypto/rand"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"dungeon-crawler/models"
)

const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

//AuthHandler ..
type AuthHandler struct {
	db mongo.Database
	googleOauthConfig *oauth2.Config
	store *sessions.CookieStore
}

//NewAuthHandler ..
func NewAuthHandler(db mongo.Database, googleOauthConfig *oauth2.Config, store *sessions.CookieStore ) *AuthHandler {
	return &AuthHandler{
		db: db,
		googleOauthConfig: googleOauthConfig,
		store: store,
	}
}

//Index returns all users
func (h *AuthHandler) Index(w http.ResponseWriter, r *http.Request) {

	collection := h.db.Collection("users")
	findOptions := options.Find()

	var users []*models.User

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	json.NewEncoder(w).Encode(users)
}

//OauthGoogleLogin ..
func (h *AuthHandler) OauthGoogleLogin(w http.ResponseWriter, r *http.Request) {
    // Create oauthState cookie
    oauthState := generateStateOauthCookie(w)
    u := h.googleOauthConfig.AuthCodeURL(oauthState)
    http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
    var expiration = time.Now().Add(365 * 24 * time.Hour)

    b := make([]byte, 16)
    rand.Read(b)
    state := base64.URLEncoding.EncodeToString(b)
    cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
    http.SetCookie(w, &cookie)

    return state
}

//OauthGoogleCallback ..
func (h *AuthHandler) OauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
    // Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")
	session, err := h.store.Get(r, "cookie-name")
	if (err != nil) {
		//todo error handling
		return
	}

    if r.FormValue("state") != oauthState.Value {
        log.Println("invalid oauth google state")
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }
	data, err := h.getUserDataFromGoogle(r.FormValue("code"))
	type userInfo struct {
		Email string `json:"email"`
		VerifiedEmail bool `json:"verifiedEmail"`
		Name string `json:"name"`
	}
	var ui userInfo
	err = json.Unmarshal(data, &ui)
    if err != nil {
        log.Println(err.Error())
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

	userCollection := h.db.Collection("users");

	var u models.User
	err = userCollection.FindOne(context.TODO(), bson.M{"email": ui.Email}).Decode(&u)
	if (err != nil) {
		u = models.User{Email: ui.Email, Name: ui.Name, VerifiedEmail: ui.VerifiedEmail}
		_, err := h.db.Collection("users").InsertOne(context.TODO(), u)
		if (err != nil) {
			//todo error handling
			return
		}
	}
	session.Values["user"] = u
    session.Values["authenticated"] = true
	err = session.Save(r, w)
	if (err != nil) {
		//todo error handling
		return
	}

	http.Redirect(w, r, "/", 302)

}

func (h *AuthHandler) getUserDataFromGoogle(code string) ([]byte, error) {
    // Use code to get token and get user info from Google.

    token, err := h.googleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
    }
    response, err := http.Get(oauthGoogleURLAPI + token.AccessToken)
    if err != nil {
        return nil, fmt.Errorf("failed getting user info: %s", err.Error())
    }
    defer response.Body.Close()
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, fmt.Errorf("failed read response: %s", err.Error())
    }
    return contents, nil
}

//CheckLoggedIn checks if user is logged in when frontend loads
func (h *AuthHandler) CheckLoggedIn(w http.ResponseWriter, r *http.Request) {
	store := h.store
	session, _ := (*store).Get(r, "cookie-name")

	if session.Values["user"] != nil && session.Values["authenticated"] == true {
		json.NewEncoder(w).Encode(session.Values["user"])
		return
	}

	json.NewEncoder(w).Encode(nil)
}
