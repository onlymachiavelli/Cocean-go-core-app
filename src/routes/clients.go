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

	// businessGroup.POST("", func(c echo.Context) error {
	// 	return handlers.CreateBusiness(c, db)
	// })
	// businessGroup.GET("", func(c echo.Context) error {
	// 	return handlers.GetMyBusinesses(c, db)
	// })
	// businessGroup.GET("/one/:id", func(c echo.Context) error {
	// 	return handlers.GetMyBusiness(c, db)
	// })

	//create order 
	cGrp.GET("/:id", func(c echo.Context) error {
		return handlers.GetMyClients(c, db)
	})
} 