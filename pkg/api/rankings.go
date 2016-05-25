package api

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/keavon/clout/pkg/game"
	"github.com/labstack/echo"
)

// Rankings returns the rankings an existing game
func (api API) Rankings(c echo.Context) error {
	// Get a connection to the redis pool
	rc := api.Pool.Get()
	defer rc.Close()

	g, err := game.Load(rc, c.Param("game"))
	if err != nil {
		if err == redis.ErrNil {
			return invalidGameIDError(c)
		}

		return err
	}

	return c.JSON(http.StatusOK, g.Rankings())
}
