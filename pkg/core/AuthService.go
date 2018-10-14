package core

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *sql.DB
}

func (authService *AuthService) Startup() {
	var err error
	authService.db, err = sql.Open("mysql", "root:admin@/fetch")
	if err != nil {
		log.Fatal(err)
	}
}

// Signup creates a new user and makes an entry in the database.
func (authService *AuthService) Signup(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	stmt, err := authService.db.Prepare("INSERT INTO users(email, password, username) VALUES(?, ?, ?)")
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
func (authService *AuthService) Signin(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := authService.db.QueryRow("SELECT password FROM users WHERE username = ?", user.Username)
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
