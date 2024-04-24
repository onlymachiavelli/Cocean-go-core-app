package handlers

import (
	"fmt"

	"cocean.com/src/requests"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("hit the business handlers package")
}



//create business 

func CreateBusiness(c echo.Context , db *gorm.DB) error{

	//bind the request body 
	payload  := new(requests.CreateBusiness)
	if err := c.Bind(payload); err != nil {
		//response 
		return c.JSON(400, map[string]interface{}{
			"message": "Invalid request payload",
		})
		
	}
	fmt.Println(payload)
	return c.JSON(200, map[string]interface{}{
		"message": "Business created successfully",
	})
}



//return ressourses 
func GetMyBusinesses (c echo.Context, db *gorm.DB) error {
	return nil 
}




//return one record 
func GetMyBusiness(c echo.Context , db *gorm.DB) error {

	return nil

}

