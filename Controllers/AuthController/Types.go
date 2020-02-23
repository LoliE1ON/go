package AuthController

import (
	"github.com/LoliE1ON/go/Models/TokenModel"
	"github.com/LoliE1ON/go/Models/UserModel"
)

type LoginRequestParams struct {
	Login    string
	Password string
}

type RegisterRequestParams struct {
	Login    string `validate:"required,max=20,min=4"`
	Name     string `validate:"required"`
	Password string `validate:"required,max=40,min=6"`
}

type ResponseSuccess struct {
	User  UserModel.User
	Token TokenModel.UserToken
}

type ResponseRegisterSuccess struct {
	Token TokenModel.UserToken
}
