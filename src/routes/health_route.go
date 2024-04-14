package route

import (
	"fmt"

	"cocean.com/src/handlers"
	"github.com/labstack/echo/v4"
)



func init() {
	fmt.Println("Hit the health go gile handler")
}



func HealthRoute(g *echo.Group) {
	g.GET("", handlers.Health)
}