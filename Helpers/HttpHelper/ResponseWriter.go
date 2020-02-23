package HttpHelper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LoliE1ON/go/Types"
)

func ResponseWriter(w http.ResponseWriter, data Types.ResponseData, code int) {

	if code == 0 {
		code = 200
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Println("Marshaling output failed:", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if _, err = w.Write(jsonBytes); err != nil {
		log.Println("Writing response failed:", err)
	}

}
