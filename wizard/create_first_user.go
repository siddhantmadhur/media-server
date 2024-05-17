package wizard

import (
	"context"
	"crypto/sha256"
	"errors"

	"ocelot/storage"
)

func createFirstUser(username string, password string, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("Password does not match")
	}
	if username == "" {
		return errors.New("Username is empty")
	}
	conn, queries, err := storage.GetConn()
	defer conn.Close()
	if err != nil {
		return err
	}
	h := sha256.New()
	hash := h.Sum([]byte(password))
	err = queries.CreateProfile(context.Background(), storage.CreateProfileParams{
		Username: username,
		Password: string(hash),
		Type:     0,
	})
	return err
}
