package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LoliE1ON/go/Controllers/IndexController"

	"github.com/gorilla/mux"

	"github.com/LoliE1ON/go/Types"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	var config = Types.ServerConfig{Port: 3500}

	router := mux.NewRouter()
	router.Use(accessControlMiddleware)
	router.HandleFunc("/", IndexController.Action).Methods("GET")

	log.Println("Server started at port:", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router); err != nil {
		log.Println("Starting server failed: ", err)
	}

}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
