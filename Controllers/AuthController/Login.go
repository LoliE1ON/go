package AuthController

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/LoliE1ON/go/Models/UserModel"

	"github.com/LoliE1ON/go/Helpers/HttpHelper"
	"github.com/LoliE1ON/go/Types"
)

func LoginAction(w http.ResponseWriter, r *http.Request) {

	var response Types.ResponseData
	var requestParams LoginRequestParams

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

	user, err := UserModel.GetByLogin(requestParams.Login)
	if err != nil {
		log.Println("LoginAction: Failed find user", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	spew.Dump(user)

	response.Data = user

	HttpHelper.ResponseWriter(response, w)
}
