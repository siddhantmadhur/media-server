package auth

import (
	"log"
	"net/http"
	"ocelot/config"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Use this middleware to authenticate routes
func AuthenticateRoute(next RouteWithUser, doOnWizard bool) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cfg config.Config
		cfg.Read()

		if !cfg.FinishedWizard && doOnWizard {
			return next(c, nil)
		}

		authorization := c.Request().Header["Authorization"]
		if len(authorization) != 1 {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "You have not included the authorization token",
			})

		}
		var config config.Config
		err := config.Read()
		if err != nil {
			return c.String(500, err.Error())
		}

		var user User
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

		claims := token.Claims.(jwt.MapClaims)
		data := claims["data"].(map[string]interface{})
		userId, err := strconv.Atoi(data["id"].(string))
		user.ID = int64(userId)
		user.Username = data["username"].(string)
		user.SessionToken = data["session"].(string)
		user.ExpiresAt = int64(claims["exp"].(float64))
		user.JwtTokenString = bearerToken
		return next(c, &user)
	}
}

type RouteWithUser func(echo.Context, *User) error
