package route

import (
	"cocean.com/src/handlers"
	"cocean.com/src/middlewares"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ClientRoute(g *echo.Group, db *gorm.DB) {
	cGrp := g.Group("/clients")

	cGrp.Use(middlewares.AuthMiddleware)

	cGrp.GET("/:id", func(c echo.Context) error {
		return handlers.GetMyClients(c, db)
	})

	cGrp.POST("/broadcast/:id", func(c echo.Context) error {
		return handlers.Broadcast(c, db)
	})
} 