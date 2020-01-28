package session

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Surafeljava/Court-Case-Management-System/rtoken"
	"github.com/dgrijalva/jwt-go"
)

// Create creates and sets session cookie
func Create(claims jwt.Claims, sessionID string, signingKey []byte, w http.ResponseWriter) {

	signedString, err := rtoken.Generate(signingKey, claims)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	c := http.Cookie{
		Name:     sessionID,
		Value:    signedString,
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
}

// Valid validates client cookie value
func Valid(cookieValue string, signingKey []byte) (bool, error) {
	valid, err := rtoken.Valid(cookieValue, signingKey)
	if err != nil || !valid {
		return false, errors.New("Invalid Session Cookie")
	}
	return true, nil
}

// Remove expires existing session
func Remove(sessionID string, w http.ResponseWriter) {
	c := http.Cookie{
		Name:    sessionID,
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
		Value:   "",
	}
	http.SetCookie(w, &c)
}
