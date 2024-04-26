package route

import (
	"fmt"

	"cocean.com/src/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)
func init() {
	fmt.Println("Hit the adminroutes.go file")
}
func AdminRoutes(g *echo.Group, db *gorm.DB) {

	g.POST("", func(c echo.Context) error {
		return handlers.CreateAdmin(c, db)
	})
	g.POST("/auth/login", func(c echo.Context) error {
		return handlers.Login(c, db)
	})
	
	
}