package TokenModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserToken struct {
	UserId primitive.ObjectID `bson:"_id" json:"userId"`
	Token  string             `json:"token,omitempty"`
	Exp    int64              `json:"exp,omitempty"`
}
