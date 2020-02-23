package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LoliE1ON/go/Helpers/ConfigHelper"
	"github.com/LoliE1ON/go/Router"

	"github.com/LoliE1ON/go/Net/Db/Mongo"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	// Get config file
	config, err := ConfigHelper.GetConfig()
	if err != nil {
		log.Println(err)
		return
	}

	// Connection to MongoDB
	err = Mongo.Connect(config.Mongo)
	if err != nil {
		err = errors.Wrap(err, "Error connection to MongoDB")
		log.Println(err)
		return
	}

	// Router
	router := mux.NewRouter()
	router.Use(Router.AccessControlMiddleware)
	Router.Routes(router)

	// Server
	log.Println("Server started at port:", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router); err != nil {
		log.Println("Starting server failed: ", err)
	}

}
