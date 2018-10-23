package auth

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Service handles Authentication and other similar functionality within the package.
type Service struct {
	db *sql.DB
}

// Startup initializes the Service object.
func (service *Service) Startup() {
	var err error
	service.db, err = sql.Open("mysql", "root:admin@/fetch")
	if err != nil {
		log.Fatal(err)
	}
}

// Signup creates a new user and makes an entry in the database.
func (service *Service) Signup(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	stmt, err := service.db.Prepare("INSERT INTO users(email, password, username) VALUES(?, ?, ?)")
	if err != nil {
		log.Print(err)
	}

	res, err := stmt.Exec(user.Email, string(hashedPassword), user.Username)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Print(res)
}

// Signin compares inputted user credentials against the user database.
func (service *Service) Signin(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := service.db.QueryRow("SELECT password FROM users WHERE username = ?", user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	storedCreds := &User{}
	err = result.Scan(&storedCreds.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(user.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
