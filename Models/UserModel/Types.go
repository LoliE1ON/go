package UserModel

type UserRole int

type User struct {
	Login    string     `json:",omitempty"`
	Password string     `json:",omitempty"`
	Role     []UserRole `json:",omitempty"`
}
