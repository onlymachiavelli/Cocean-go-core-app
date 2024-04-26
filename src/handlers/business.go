package handlers

import (
	"fmt"

	"cocean.com/src/models"
	"cocean.com/src/requests"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("hit the business handlers package")
}

func CreateBusiness(c echo.Context , db *gorm.DB) error{

	//bind the request body 
	payload  := new(requests.CreateBusiness)
	if err := c.Bind(payload); err != nil {
		//response 
		return c.JSON(400, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	owner := c.Get("Owner")

	fmt.Println("The payload is : " , payload)
	if payload.Address == "" || payload.Name == "" || payload.Phone == ""  || payload.Email == "" || payload.Description == "" || payload.Photo == ""	 {
		return c.JSON(400, map[string]interface{}{
			"message": "All fields are required",
		})
	}

	business := models.Business{
		Name: payload.Name,
		Address: payload.Address,
		Phone: payload.Phone,
		Email: payload.Email,
		Description: payload.Description,
		Owner:  owner.(int),
	}

	err := db.Create(&business).Error
	if (err != nil) {
		return c.JSON(500, map[string]interface{}{
			"message": "Could not create business",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "Business created successfully",
		"business": business,
	})


}



//return ressourses 
func GetMyBusinesses (c echo.Context, db *gorm.DB) error {

	id := c.Get("Owner").(int)
	if id == 0 {
		return c.JSON(400, map[string]interface{}{
			"message": "Invalid user",
		})
	}

	var businesses []models.Business
	err := db.Where("owner = ?", id).Find(&businesses).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": "Could not fetch businesses",
		})
	}
	return c.JSON(200, map[string]interface{}{
			"businesses": businesses,
		})

}
