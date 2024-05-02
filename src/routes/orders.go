package route

import (
	"cocean.com/src/handlers"
	"cocean.com/src/middlewares"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func OrderRoute(g *echo.Group, db *gorm.DB) {
	orderGrp := g.Group("/order")

	orderGrp.Use(middlewares.AuthMiddleware)

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
	orderGrp.POST("/:id", func(c echo.Context) error {
		return handlers.CreateOrder(c, db)
	})

	orderGrp.GET("/:id", func(c echo.Context) error {
		return handlers.GetOrders(c, db)
	})
} 