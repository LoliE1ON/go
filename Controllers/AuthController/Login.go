package AuthController

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

	// Select user
	user, err := UserModel.GetByLogin(requestParams.Login)
	if err != nil {
		response.Error = "User not found"
		HttpHelper.ResponseWriter(w, response, http.StatusForbidden)
		return
	}

	// Check password
	if user.Password != passwordToHash(requestParams.Password) {
		response.Error = "Login or password incorrect"
		HttpHelper.ResponseWriter(w, response, http.StatusForbidden)
		return
	}

	response.Data = user
	HttpHelper.ResponseWriter(w, response, http.StatusOK)

}

// Convert password to MD5 Hash
// md5 -_-
func passwordToHash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
