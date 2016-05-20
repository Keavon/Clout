package api

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/keavon/clout/pkg/authorization"
	"github.com/keavon/clout/pkg/country"
	"github.com/keavon/clout/pkg/game"
	"github.com/keavon/clout/pkg/token"
	"github.com/labstack/echo"
)

// joinRequest is the body for a create request
type joinRequest struct {
	GameID   string `json:"gameid"`
	Username string `json:"username"`
}

// joinResponse is the body for a create response
type joinResponse struct {
	Token   string          `json:"token"`
	Country country.Country `json:"country"`
}

// Join joins an existing game
func (api API) Join(c echo.Context) error {
	req := new(joinRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if req.Username == "" {
		return invalidUsernameError(c)
	}

	if req.GameID == "" {
		return invalidGameIDError(c)
	}

	// Get a connection to the redis pool
	rc := api.Pool.Get()
	defer rc.Close()

	g, err := game.Load(rc, req.GameID)
	if err != nil {
		if err == redis.ErrNil {
			return invalidGameIDError(c)
		}

		return err
	}

	// Create new player
	playerID, err := token.Token()
	if err != nil {
		return err
	}

	player, err := g.NewPlayer(playerID, req.Username, false)
	if err != nil {
		if err == game.ErrFull {
			return gameFullError(c)
		}

		return err
	}

	// Save game and player to redis
	err = g.Save(rc)

	if err != nil {
		return err
	}

	err = player.Save(rc)

	if err != nil {
		return err
	}

	// Create authorization for player
	token, err := token.Token()
	if err != nil {
		return err
	}
	auth := authorization.New(token, g, player)

	// Save authorization to redis
	err = auth.Save(rc)

	if err != nil {
		return err
	}

	res := joinResponse{Token: token, Country: player.Country}

	return c.JSON(http.StatusCreated, res)
}
