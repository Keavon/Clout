package api

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/keavon/clout/pkg/authorization"
	"github.com/labstack/echo"
)

// Player gets information about current player
func (api API) Player(c echo.Context) error {
	token := c.Request().Header().Get("Authorization")

	if token == "" {
		return invalidTokenError(c)
	}

	// Get a connection to the redis pool
	rc := api.Pool.Get()
	defer rc.Close()

	auth, err := authorization.Load(rc, token)
	if err != nil {
		if err == redis.ErrNil {
			return invalidTokenError(c)
		}

		return err
	}

	err = auth.Touch(rc)
	if err != nil {
		return err
	}

	err = auth.Player.Touch(rc)
	if err != nil {
		return err
	}

	err = auth.Game.Touch(rc)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, auth.Player)
}
