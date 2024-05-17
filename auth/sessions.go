package auth

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"ocelot/storage"
	"time"
)

func createSession(userId int, deviceName string, clientName string, clientVersion string, device string) (storage.Session, error) {
	var session storage.Session
	accessToken, err := rsa.GenerateKey(rand.Reader, 16*8)
	if err != nil {
		return session, err
	}

	conn, queries, err := storage.GetConn()
	defer conn.Close()

	session, err = queries.CreateSession(context.Background(), storage.CreateSessionParams{
		ID:            accessToken.N.Text(62),
		UserID:        int64(userId),
		Device:        device,
		DeviceName:    deviceName,
		ClientName:    clientName,
		ClientVersion: clientVersion,
		CreatedAt:     time.Now(),
		ExpiresAt:     time.Now().Add(time.Hour * 744),
	})

	return session, err
}
