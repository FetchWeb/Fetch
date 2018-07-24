package core

// Auth struct contains functions to authenticate
type Auth struct {
}

// User struct to store
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	password string
}

// ApiToken stores the api_token JSON response object
type ApiToken struct {
	UserID    string `json:"user_id"`
	AuthToken string `json:"auth_token"`
	Expiry    string `json:"expiry"`
}

func (auth *Auth) login(username string, password string) (*User, error) {
	return &User{}, nil
}
