package library

import (
	"fmt"
	"ocelot/auth"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

// /server/information/folders?directory=/users/...
func GetPathFolders(c echo.Context, u *auth.User) error {
	currentPath := c.QueryParam("directory")
	if len(currentPath) == 0 {
		currentPath = "/"
	}

	dir, err := os.ReadDir(currentPath)
	if err != nil {
		return c.String(500, fmt.Sprintf("Could not read path. %s", err.Error()))
	}
	var results = []map[string]string{}
	for _, entry := range dir {
		if entry.IsDir() && strings.Index(entry.Name(), ".") != 0 {
			results = append(results, map[string]string{
				"name": entry.Name(),
			})
		}
	}
	return c.JSON(200, results)
}
