package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Im Alive Bitch")
}