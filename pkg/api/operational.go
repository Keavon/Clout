package api

import (
	"net/http"
	"strconv"

	"github.com/keavon/clout/pkg/authorization"
	"github.com/keavon/clout/pkg/game"
	"github.com/keavon/clout/pkg/player"
	"github.com/keavon/clout/pkg/resource"
	"github.com/labstack/echo"
)

// createRequest is the body for a create request
type operationalRequest struct {
	Number int `json:"number"`
}

// Operational sets the number of operational instances
func (api API) Operational(c echo.Context) error {
	auth := c.Get("auth").(authorization.Authorization)

	if auth.Game.Status == game.Stopped {
		return gameEndedError(c)
	}

	req := new(operationalRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	id := c.Param("id")

	if id == "" {
		return invalidResourceError(c)
	}

	res, err := strconv.Atoi(id)

	if err != nil {
		return invalidResourceError(c)
	}

	err = auth.Player.SetOperational(res, req.Number)

	if err != nil {
		// Handle possible errors from setting operational instances
		if err == resource.ErrInvalidResource {
			return invalidResourceError(c)
		}

		if err == player.ErrInvalidOperational {
			return operationalError(c)
		}
	}

	// Get a connection to the redis pool
	rc := api.Pool.Get()
	defer rc.Close()

	err = auth.Player.Save(rc)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
