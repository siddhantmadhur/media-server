package auth

import "github.com/labstack/echo/v4"

// Return jwt on success
func Login(c echo.Context) error {
	var response struct {
		Username      string `json:"username"`
		Password      string `json:"password"`
		Device        string `json:"device"`
		DeviceName    string `json:"deviceName"`
		ClientName    string `json:"clientName"`
		ClientVersion string `json:"clientVersion"`
	}

	err := c.Bind(&response)
	if err != nil {
		return c.JSON(500, err)
	}
	var user User
	err = user.Login(response.Username, response.Password, response.Device, response.DeviceName, response.ClientName, response.ClientVersion)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, map[string]string{
		"token": user.JwtTokenString,
	})
}
