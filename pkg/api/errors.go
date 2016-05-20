package api

import (
	"net/http"

	"github.com/labstack/echo"
)

type errorCode struct {
	ErrorID int `json:"errorId"`
}

func invalidUsernameError(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, errorCode{ErrorID: 0})
}

func invalidTokenError(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, errorCode{ErrorID: 1})
}

func invalidGameIDError(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, errorCode{ErrorID: 2})
}

func gameFullError(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, errorCode{ErrorID: 3})
}
