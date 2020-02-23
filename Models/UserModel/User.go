package UserModel

import (
	"context"
	"log"

	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/LoliE1ON/go/Net/Db/Mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll() (users []User, err error) {

	collection := Mongo.GetDatabase().Collection("users")

	opts := options.Find().SetProjection(bson.D{
		{"password", 0},
	})

	cur, err := collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		err = errors.Wrap(err, "Net error! Find error of UserModel.GetAll")
		return
	}

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		var user User

		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}
	if err = cur.Err(); err != nil {
		return
	}

	return

}

func GetByLogin(login string) (user User, err error) {

	collection := Mongo.GetDatabase().Collection("users")
	opts := options.FindOne().SetProjection(bson.D{
		{"role", 0},
	})

	err = collection.FindOne(context.TODO(), bson.M{"login": login}, opts).Decode(&user)
	if err != nil {
		err = errors.Wrap(err, "Net error! Error select user: UserModel.GetByLogin")
		return
	}

	return
}
