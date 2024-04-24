package middlewares

import (
	"fmt"

	"cocean.com/src/utils"
	"github.com/labstack/echo/v4"
)


func init() {
	fmt.Println("Hit the auth middleware")
}


func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Hit the auth middleware")
		//get the token from the header 
		//token bearer 
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(400, map[string]interface{}{
				"message": "Token is required",
			})
		}

		token= token[7:]
		//validate the token 
		claims, err := utils.VerifyToken(token)
		if err != nil {
			return c.JSON(400, map[string]interface{}{
				"message": "Invalid token",
			})
		}

		OnwerId := claims["id"]

		
		

		c.Set("Owner", OnwerId)
		return next(c)

	
	}
}