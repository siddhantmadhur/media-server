package auth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"ocelot/config"
	"ocelot/storage"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func createUser(username string, password string, confirmPassword string, priv int64) error {
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
	h.Write([]byte(password))
	hash := base64.URLEncoding.EncodeToString(h.Sum(nil))
	err = queries.CreateProfile(context.Background(), storage.CreateProfileParams{
		Username: username,
		Password: string(hash),
		Type:     priv,
	})
	return err
}

// Creates user with admin rights if server is unitialized, otherwise creates regular user
func CreateUser(c echo.Context) error {
	var request struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	err := c.Bind(&request)
	if err != nil {
		return c.String(500, err.Error())
	}

	var config config.Config
	err = config.Read()
	if err != nil {
		return c.String(500, err.Error())
	}
	if config.FinishedWizard {
		authorization := c.Request().Header["Authorization"]
		if len(authorization) != 1 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "You have not included the authorization token",
			})
		}
		err := config.Read()
		if err != nil {
			return c.String(500, err.Error())
		}
		bearerToken := strings.ReplaceAll(authorization[0], "Bearer ", "")
		token, err := jwt.Parse(bearerToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.SecretKey), nil
		})
		if err != nil {
			log.Printf("[ERROR]: (auth/authenticate.go) %s\n", err.Error())
			return c.String(500, err.Error())
		}
		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "Token not valid",
			})
		}

		err = createUser(request.Username, request.Password, request.ConfirmPassword, 1)
		if err != nil {
			return c.String(500, err.Error())
		}
		return c.NoContent(201)

	} else {
		conn, queries, err := storage.GetConn()
		defer conn.Close()
		if err != nil {
			return c.String(500, err.Error())
		}
		_, err = queries.GetAdminUser(context.Background())

		// err != nil implies there is no admin user
		if err != nil {
			err = createUser(request.Username, request.Password, request.ConfirmPassword, 0)
			if err != nil {
				return c.String(500, err.Error())
			}
		} else {
			if request.Password != request.ConfirmPassword {
				err = errors.New("Password does not match")
				return c.String(500, err.Error())
			}
			if len(request.Username) == 0 {
				err = errors.New("Username is empty")
				return c.String(500, err.Error())
			}
			conn, queries, err := storage.GetConn()
			defer conn.Close()
			if err != nil {
				return err
			}
			h := sha256.New()
			h.Write([]byte(request.Password))
			hash := base64.URLEncoding.EncodeToString(h.Sum(nil))
			err = queries.UpdateAdminUser(context.Background(), storage.UpdateAdminUserParams{
				Username: request.Username,
				Password: string(hash),
			})
			if err != nil {
				return c.String(500, err.Error())
			}
		}
		return c.NoContent(201)
	}
}
