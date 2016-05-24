package api

import (
	"net/http"

	"github.com/keavon/clout/pkg/authorization"
	"github.com/keavon/clout/pkg/game"
	"github.com/labstack/echo"
)

// Start starts a game
func (api API) Start(c echo.Context) error {
	auth := c.Get("auth").(authorization.Authorization)

	if !auth.Player.Admin {
		return adminPermsRequiredError(c)
	}

	if auth.Game.Status != game.Waiting {
		return gameAlreadyStartedError(c)
	}

	auth.Game.Start()

	// Get a connection to the redis pool
	rc := api.Pool.Get()
	defer rc.Close()

	err := auth.Game.Save(rc)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
