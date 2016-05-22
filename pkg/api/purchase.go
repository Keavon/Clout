package api

import (
	"net/http"
	"strconv"

	"github.com/keavon/clout/pkg/authorization"
	"github.com/keavon/clout/pkg/player"
	"github.com/keavon/clout/pkg/resource"
	"github.com/labstack/echo"
)

// Purchase handles installation purchases
func (api API) Purchase(c echo.Context) error {
	auth := c.Get("auth").(authorization.Authorization)

	id := c.Param("id")

	if id == "" {
		return invalidResourceError(c)
	}

	res, err := strconv.Atoi(id)

	if err != nil {
		return invalidResourceError(c)
	}

	err = auth.Player.Purchase(res)

	if err != nil {
		// Handle possible errors from purchasing
		if err == resource.ErrInvalidResource {
			return invalidResourceError(c)
		}

		if err == player.ErrCost {
			return costExceededError(c)
		}

		if err == player.ErrCapacity {
			return capacityError(c)
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
