package AuthController

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LoliE1ON/go/Models/TokenModel"

	"github.com/LoliE1ON/go/Strategies/JwtStrategy"

	"github.com/LoliE1ON/go/Models/UserModel"

	"github.com/LoliE1ON/go/Helpers/HttpHelper"
	"github.com/LoliE1ON/go/Types"
)

// User login
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

	// Create new token
	token, err := JwtStrategy.CreateToken(user.UserId)
	if err != nil {
		log.Println("Failed to create JW Token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	insertToken := TokenModel.UserToken{
		UserId: user.UserId,
		Token:  token,
	}

	// Save token to db
	err = TokenModel.Insert(&insertToken)
	if err != nil {
		log.Println("Failed to insert JW Token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	data := ResponseSuccess{
		User:  user,
		Token: insertToken,
	}

	response.Data = data
	HttpHelper.ResponseWriter(w, response, http.StatusOK)

}

// Convert password to MD5 Hash
// TODO: Replace to Bcrypt, create helper
func passwordToHash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
