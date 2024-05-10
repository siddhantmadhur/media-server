package config

import (
	"context"
	"crypto/sha256"
	"errors"

	"github.com/siddhantmadhur/media-server/storage"
)

// Create the first admin user for the server
func createAdminUser(username string, password string, confirmPassword string) error {
	if password != confirmPassword {
		return errors.New("Password does not match")
	}
	conn, queries, err := storage.GetConn()
	defer conn.Close()
	if err != nil {
		return err
	}
	is_finished_setup, err := queries.IsFinishedSetup(context.Background())
	if err != nil {
		return err
	}
	if is_finished_setup > 0 {
		return errors.New("Admin user already present")
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
