package Router

import (
	"github.com/LoliE1ON/go/Controllers/AuthController"
	"github.com/LoliE1ON/go/Controllers/IndexController"
	"github.com/gorilla/mux"
)

// Routes list
func Routes(router *mux.Router) {

	router.HandleFunc("/", IndexController.Action).Methods("GET")
	router.HandleFunc("/auth/login", AuthController.LoginAction).Methods("POST")

}
