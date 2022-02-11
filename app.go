////////////////////////////////////////////////////////////////////////////////
//	app.go  -  Feb/10/2022  -  aldebap
//
//	personLib App
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	databaseURL string
	dbClient    *mongo.Client
	httpRouter  *mux.Router
}

//	create a new application
func New(databaseURL string) *App {

	return &App{
		databaseURL: databaseURL,
	}
}

//	run the application
func (a *App) Run(portNumber string) {

	var err error

	//	connect to the database
	clientOptions := options.Client().ApplyURI(a.databaseURL)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	a.dbClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer a.dbClient.Disconnect(ctx)

	//	start the Web Server
	a.httpRouter = mux.NewRouter()

	//	a.Router.HandleFunc("/author", author.getAll).Methods("GET")
	//	a.Router.HandleFunc("/publisher", publisher.getAll).Methods("GET")
	http.Handle("/", a.httpRouter)

	//start and listen to requests
	fmt.Printf("Listening port %s\n", portNumber)

	log.Panic(http.ListenAndServe(portNumber, a.httpRouter))
}
