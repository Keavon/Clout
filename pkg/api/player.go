package api

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/keavon/clout/pkg/player"
	"github.com/labstack/echo"
)

// Player gets information about current player
func (a API) Player(c echo.Context) error {
	token := c.Request().Header().Get("Authorization")

	if token == "" {
		return invalidTokenError(c)
	}

	// Get a connection to the redis pool
	rc := a.Pool.Get()
	defer rc.Close()

	p, err := player.Load(rc, token)
	if err != nil {
		if err == redis.ErrNil {
			return invalidTokenError(c)
		}

		return err
	}

	err = p.Touch(rc)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}
