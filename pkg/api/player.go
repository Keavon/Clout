package api

import (
	"net/http"

	"github.com/keavon/clout/pkg/authorization"
	"github.com/keavon/clout/pkg/game"
	"github.com/labstack/echo"
)

// Player gets information about current player
func (api API) Player(c echo.Context) error {
	auth := c.Get("auth").(authorization.Authorization)

	if auth.Game.Status == game.Started {
		// First check the player will not have their update time set
		// Set it to when the game Started
		if auth.Player.LastUpdated.IsZero() {
			auth.Player.LastUpdated = auth.Game.StartTime
		}
		auth.Player.Update(auth.Game.StartTime)

		// Get a connection to the redis pool
		rc := api.Pool.Get()
		defer rc.Close()

		err := auth.Player.Save(rc)
		if err != nil {
			return err
		}
	}

	// TODO: Quick ugly hack, fix logic later
	if auth.Game.Status == game.Waiting {
		auth.Player.Demand = auth.Player.Country.InitialDemand
	}

	return c.JSON(http.StatusOK, auth)
}
