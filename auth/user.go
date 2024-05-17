package auth

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"ocelot/storage"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	SessionToken string `json:"session_token"`
	ExpiresAt    string `json:"expires_at"`
}

func (u *User) Login(username string, password string) error {
	conn, queries, err := storage.GetConn()
	defer conn.Close()
	if err != nil {
		return err
	}

	h := sha256.New()

	hashedPass := h.Sum([]byte(password))

	user, err := queries.GetUserWithPassword(context.Background(), storage.GetUserWithPasswordParams{
		Username: username,
		Password: string(hashedPass),
	})
	fmt.Printf("user: %s \n", user.Username)

	if err != nil {
		return errors.New("No user found.")
	}

	var data struct {
		Id           string    `json:"id"`
		Username     string    `json:"username"`
		Type         int       `json:"type"`
		ExpiresAt    time.Time `json:"expires_at"`
		CreatedAt    time.Time `json:"created_at"`
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
	}

	// TODO
	// Change this to an environment variable
	//var secret = "this-is-a-secret-32"
	token := jwt.New(jwt.SigningMethodEdDSA)
	fmt.Printf("data: %s\n", data.AccessToken)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(60 * time.Minute)
	claims["username"] = username
	claims["username"] = username
	claims["username"] = username

	return nil
}
