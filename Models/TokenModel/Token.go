package TokenModel

import (
	"context"
	"log"
	"time"

	"github.com/LoliE1ON/go/Helpers/ConfigHelper"

	"github.com/pkg/errors"

	"github.com/LoliE1ON/go/Net/Db/Mongo"
	"go.mongodb.org/mongo-driver/bson"
)

// Select token
func Select(token string) (userToken UserToken, err error) {

	collection := Mongo.GetDatabase().Collection("tokens")

	err = collection.FindOne(context.TODO(), bson.M{"token": token}).Decode(&userToken)
	if err != nil {
		err = errors.Wrap(err, "Net error! Error select token")
		return
	}

	return
}

// Insert token to db
func Insert(userToken *UserToken) (err error) {

	// Get config file
	config, err := ConfigHelper.GetConfig()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(time.Now().Unix() + (config.JwtExpHours * 3600))
	userToken.Exp = time.Now().Unix() + (config.JwtExpHours * 3600)

	collection := Mongo.GetDatabase().Collection("tokens")

	_, err = collection.InsertOne(context.Background(),
		bson.D{
			{"userId", userToken.UserId},
			{"token", userToken.Token},
			{"exp", userToken.Exp},
		})
	if err != nil {
		err = errors.Wrap(err, "Net error! Error insert token")
		return
	}

	return
}
