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
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"personaLib/author"
)

type Config struct {
	databaseURL string
	timeout     time.Duration
	servicePort string
}

type App struct {
	config     *Config
	dbClient   *mongo.Client
	httpRouter *mux.Router
}

//	create a new application
func GetFromEnv() *Config {

	//	get configuration parameters from environment
	databaseURL := os.Getenv("DATABASEURL")
	servicePort := os.Getenv("SERVICEPORT")

	if len(servicePort) == 0 {
		servicePort = ":8080"
	} else if servicePort[0] != ':' {
		servicePort = ":" + servicePort
	}

	return &Config{
		databaseURL: databaseURL,
		timeout:     10 * time.Second,
		servicePort: servicePort,
	}
}

//	run the application
func (a *App) Run(config *Config) {

	var err error

	a.config = config

	//	connect to the database
	clientOptions := options.Client().ApplyURI(a.config.databaseURL)
	ctx, cancel := context.WithTimeout(context.Background(), a.config.timeout)

	defer cancel()

	a.dbClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer a.dbClient.Disconnect(ctx)

	//	initialize enetities collections
	author.InitCollection(a.dbClient)

	//	start the Web Server
	a.httpRouter = mux.NewRouter()

	a.httpRouter.HandleFunc("/author", getAllAccounts).Methods("GET")
	//	a.Router.HandleFunc("/publisher", publisher.getAll).Methods("GET")
	http.Handle("/", a.httpRouter)

	//start and listen to requests
	fmt.Printf("Listening port %s\n", a.config.servicePort)

	log.Panic(http.ListenAndServe(a.config.servicePort, a.httpRouter))
}
