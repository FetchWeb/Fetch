package auth

import "github.com/FetchWeb/Fetch/pkg/core"

// User is the struct that stores credentials for authentication.
type User struct {
	core.DBObject
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Username string `json:"username" db:"username"`
}
