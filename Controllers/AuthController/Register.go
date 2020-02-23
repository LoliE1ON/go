package AuthController

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LoliE1ON/go/Helpers/ValidateHelper"

	"github.com/LoliE1ON/go/Helpers/HttpHelper"
	"github.com/LoliE1ON/go/Types"
)

// User register
func RegisterAction(w http.ResponseWriter, r *http.Request) {

	var response Types.ResponseData
	var requestParams RegisterRequestParams

	var requestBody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Reading request body failed:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(requestBody, &requestParams); err != nil {
		log.Println("Unmarshal request body failed:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Validate register data
	validateErrors, err := ValidateHelper.Validate(requestParams)
	if err != nil {
		log.Println("Validator error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}

	// Response validate errors
	if len(validateErrors) > 0 {
		response.Error = validateErrors
		HttpHelper.ResponseWriter(w, response, http.StatusBadRequest)
		return
	}

	/*data := ResponseSuccess{
		User:  user,
		Token: insertToken,
	}*/

	response.Data = "Register"
	HttpHelper.ResponseWriter(w, response, http.StatusOK)

}
