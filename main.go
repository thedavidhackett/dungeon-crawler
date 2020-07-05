package main

import (
	"context"
	"fmt"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"dungeon-crawler/controllers"
	"dungeon-crawler/middlewares"
	"dungeon-crawler/models"


)

var store *sessions.CookieStore

func init(){
	if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
	}

	store = sessions.NewCookieStore(
		[]byte("secret-key"),
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 30,
		HttpOnly: true,
	}

	gob.Register(models.User{})

}

func main(){

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	var googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	db := client.Database("dungeoncrawlertest")

	ah := controllers.NewAuthHandler(*db, googleOauthConfig, store)

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	auth := r.PathPrefix("/auth").Subrouter()

	api.HandleFunc("/users", ah.Index)
	auth.HandleFunc("/google/login", ah.OauthGoogleLogin)
	auth.HandleFunc("/google/callback", ah.OauthGoogleCallback)
	auth.HandleFunc("/check-logged-in", ah.CheckLoggedIn)

	spa := spaHandler{staticPath: "public", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	fmt.Println("Listening on port 8000")
	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8000",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
	}

	sm := middlewares.NewSessionMiddleware(store)
	api.Use(sm.Middleware)

    log.Fatal(srv.ListenAndServe())
}

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
        // if we failed to get the absolute path respond with a 400 bad request
        // and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

    // check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
        // if we got an error (that wasn't that the file doesn't exist) stating the
        // file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

    // otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
