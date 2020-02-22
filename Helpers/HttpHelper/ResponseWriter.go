package HttpHelper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LoliE1ON/go/Types"
)

func ResponseWriter(data Types.ResponseData, w http.ResponseWriter) {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Println("Marshaling output failed:", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(jsonBytes); err != nil {
		log.Println("Writing response failed:", err)
	}

}
