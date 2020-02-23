package JwtStrategy

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/LoliE1ON/go/Helpers/ConfigHelper"
	"github.com/dgrijalva/jwt-go"
)

// Create a new token object
func CreateToken(userId primitive.ObjectID) (tokenString string, err error) {

	// Get config file
	config, err := ConfigHelper.GetConfig()
	if err != nil {
		log.Println(err)
		return
	}

	type MyCustomClaims struct {
		UserId primitive.ObjectID `json:"userId"`
		jwt.StandardClaims
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + (config.JwtExpHours * 3600),
		},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		log.Println(err)
		return
	}

	return
}
