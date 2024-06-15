package auth

import (
	"context"
	"net/http"
	"ocelot/config"
	"ocelot/storage"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// ADMIN FUNCTION: DO NOT USE DIRECTLY. Ensure user has correct rights before using.
func updateUser(userId int64, username string, password string, queries *storage.Queries) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 24)
	if err != nil {
		return err
	}

	err = queries.UpdateUser(context.Background(), storage.UpdateUserParams{
		Username: username,
		Password: string(hashedPassword),
		ID:       userId,
	})

	return err
}

func createUser(username string, password string, queries *storage.Queries, perm int) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 24)
	if err != nil {
		return err
	}

	err = queries.CreateProfile(context.Background(), storage.CreateProfileParams{
		Username: username,
		Password: string(hashedPassword),
		Type:     int64(perm),
	})

	return err
}

func CreateNewUserRoute(c echo.Context, u *User, cfg *config.Config) error {

	// Check if the user is logged in or if the server is in wizard mode
	if u == nil && cfg.FinishedWizard {
		var result = map[string]string{
			"message": "You are not authorized",
		}
		c.Response().Header().Add("WWW-Authenticate", "Bearer token=\"Access token provided by server\"")
		return c.JSON(401, result)
	}

	conn, query, err := storage.GetConn()
	defer conn.Close()
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, result)
	}

	var request struct {
		Username      string `json:"username"`
		Password      string `json:"password"`
		PermissionInt int64  `json:"permission_int"`
	}
	err = c.Bind(&request)
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, result)
	}

	// If server is in wizard mode, no more than one account should be created
	if !cfg.FinishedWizard {
		profiles, err := query.GetProfiles(context.Background())
		if err != nil {
			var result = map[string]string{
				"message": "There was an error",
				"error":   err.Error(),
			}
			return c.JSON(http.StatusInternalServerError, result)
		}
		if len(profiles) > 0 {
			admin, err := query.GetAdminUser(context.Background())
			if err != nil {
				var result = map[string]string{
					"message": "There was an error",
					"error":   err.Error(),
				}
				return c.JSON(http.StatusInternalServerError, result)
			}
			err = updateUser(admin.ID, request.Username, request.Password, query)
			if err != nil {
				var result = map[string]string{
					"message": "There was an error",
					"error":   err.Error(),
				}
				return c.JSON(http.StatusInternalServerError, result)
			}
			return c.NoContent(200)
		}

		// if in wizard mode, ensure the new users permission int is 0
		request.PermissionInt = 0
	} else {
		// Only the server admin can create new users
		if u.PermissionLevel != 0 {
			var result = map[string]string{
				"message": "You do not have permission to create new accounts",
			}
			return c.JSON(403, result)
		}
	}

	// Create new user
	err = createUser(request.Username, request.Password, query, int(request.PermissionInt))
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, result)
	}

	return c.NoContent(201)
}
