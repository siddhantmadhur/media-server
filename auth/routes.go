package auth

import "github.com/labstack/echo/v4"

// Return jwt on success
func Login(c echo.Context) error {
	return c.String(200, "Hello, Login!")
}
