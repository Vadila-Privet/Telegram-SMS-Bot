package main

import (
	"fmt"
	"log"
	"net/http"

	route "./http"
	"./models"

	"github.com/globalsign/mgo"

	"github.com/gorilla/mux"
)

const (
	databaseURL = "mongo:27017"

	databaseName = "info"
)

func main() {

	//Create db session
	db, err := mgo.Dial(databaseURL)
	//Sends session variable to the repository
	models.Connection.Session = db
	if err != nil {
		fmt.Printf("Error is: %s", err)
	}
	defer db.Close()
	db.SetMode(mgo.Monotonic, true)

	fmt.Println("You are inside")

	r := mux.NewRouter()
	route.New(r)
	log.Fatal(http.ListenAndServe(":8181", r))
}
