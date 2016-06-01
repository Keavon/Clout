package api

import (
	"github.com/garyburd/redigo/redis"
	"github.com/keavon/clout/pkg/authorization"
	"github.com/labstack/echo"
)

// Auth is echo middleware that performs user authentication
func (api API) Auth() echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
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

			auth.Game.Update()
			err = auth.Game.Touch(rc)
			if err != nil {
				return err
			}

			c.Set("auth", auth)
			return next(c)
		}
	}
}
