package api

import (
	"net/http"

	"github.com/keavon/clout/pkg/authorization"
	"github.com/labstack/echo"
)

// Player gets information about current player
func (api API) Player(c echo.Context) error {
	auth := c.Get("auth").(authorization.Authorization)

	return c.JSON(http.StatusOK, auth.Player)
}
