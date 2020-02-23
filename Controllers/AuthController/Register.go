package AuthController

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LoliE1ON/go/Models/TokenModel"
	"github.com/LoliE1ON/go/Strategies/JwtStrategy"

	"github.com/LoliE1ON/go/Models/UserModel"

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

	// Check login exist
	_, err = UserModel.GetByLogin(requestParams.Login)
	if err == nil {
		response.Error = "Login exist"
		HttpHelper.ResponseWriter(w, response, http.StatusBadRequest)
		return
	}

	// Register new user
	userId, err := UserModel.Register(UserModel.User{
		Login:    requestParams.Login,
		Password: requestParams.Password,
		Name:     requestParams.Name,
	})
	if err != nil {
		log.Println("Net error:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}

	// Create new token
	token, err := JwtStrategy.CreateToken(userId)
	if err != nil {
		log.Println("Failed to create JW Token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Save token to db
	insertToken := TokenModel.UserToken{
		UserId: userId,
		Token:  token,
	}
	err = TokenModel.Insert(&insertToken)
	if err != nil {
		log.Println("Failed to insert JW Token:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	data := ResponseRegisterSuccess{
		Token: insertToken,
	}

	response.Data = data
	HttpHelper.ResponseWriter(w, response, http.StatusOK)

}
