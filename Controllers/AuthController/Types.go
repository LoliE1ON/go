package AuthController

import (
	"github.com/LoliE1ON/go/Models/TokenModel"
	"github.com/LoliE1ON/go/Models/UserModel"
)

type LoginRequestParams struct {
	Login    string
	Password string
}

type ResponseSuccess struct {
	User  UserModel.User
	Token TokenModel.UserToken
}
