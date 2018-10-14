package core

// User is the struct that stores credentials for authentication.
type User struct {
	DBObject
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
}
