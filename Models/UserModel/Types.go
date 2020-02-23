package UserModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRole int

type User struct {
	UserId   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Login    string             `json:"login,omitempty"`
	Password string             `json:"password,omitempty"`
	Role     []UserRole         `json:"role,omitempty"`
}
