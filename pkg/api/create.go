package api

import (
	"net/http"

	"github.com/keavon/clout/pkg/authorization"
	"github.com/keavon/clout/pkg/country"
	"github.com/keavon/clout/pkg/game"
	"github.com/keavon/clout/pkg/token"
	"github.com/labstack/echo"
)

// createRequest is the body for a create request
type createRequest struct {
	Username string `json:"username"`
}

// createResponse is the body for a create response
type createResponse struct {
	GameID  string          `json:"gameid"`
	Token   string          `json:"token"`
	Country country.Country `json:"country"`
}

// Create creates a new game
func (api API) Create(c echo.Context) error {
	req := new(createRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if req.Username == "" {
		return invalidUsernameError(c)
	}

	// Create new game
	gameid, err := token.ShortToken()
	if err != nil {
		return err
	}

	g := game.New(gameid)

	// Create new player
	playerID, err := token.Token()
	if err != nil {
		return err
	}

	player, err := g.NewPlayer(playerID, req.Username, true)
	if err != nil {
		return err
	}

	// Get a connection to the redis pool
	rc := api.Pool.Get()
	defer rc.Close()

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

	res := createResponse{GameID: gameid, Token: token, Country: player.Country}

	return c.JSON(http.StatusCreated, res)
}
