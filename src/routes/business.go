package route

import (
	"cocean.com/src/handlers"
	"cocean.com/src/middlewares"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BusinessRoutes(g *echo.Group, db *gorm.DB) {
	businessGroup := g.Group("/business")

	businessGroup.Use(middlewares.AuthMiddleware)

	businessGroup.POST("", func(c echo.Context) error {
		return handlers.CreateBusiness(c, db)
	})
	businessGroup.GET("", func(c echo.Context) error {
		return handlers.GetMyBusinesses(c, db)
	})
}