package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/LoliE1ON/go/Router"

	"github.com/LoliE1ON/go/Net/Db/MongoDb"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"

	"github.com/LoliE1ON/go/Types"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	// Parsing config file
	config, err := parseConfig()
	if err != nil {
		log.Println(err)
		return
	}

	// Connection to MongoDB
	err = MongoDb.Connect(config.Mongo)
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

// Parse config file
func parseConfig() (config Types.ServerConfig, err error) {

	const configFilePath = "serverConfig.json"

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		jsonBytes, err := json.MarshalIndent(config, "", "\t")
		if err != nil {
			return config, errors.Wrap(err, fmt.Sprintf("Marshaling config error (%s)", configFilePath))
		}

		if err := ioutil.WriteFile(configFilePath, jsonBytes, 0644); err != nil {
			return config, errors.Wrap(err, fmt.Sprintf("Write in config data file error (%s)", configFilePath))
		}
	} else {
		jsonData, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			return config, errors.Wrap(err, fmt.Sprintf("Read config file error (%s)", configFilePath))
		}

		if err := json.Unmarshal(jsonData, &config); err != nil {
			return config, errors.Wrap(err, fmt.Sprintf("Unmarshal config file error (%s)", configFilePath))
		}
	}

	return
}
