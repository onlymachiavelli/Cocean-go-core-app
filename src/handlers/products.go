package handlers

import (
	"strconv"

	"cocean.com/src/models"
	"cocean.com/src/requests"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateProduct(c echo.Context, db *gorm.DB) error {

	idBus := c.Param("id") 
	if idBus == "" {
		return c.JSON(400 , "Missing parameter: bus_id")
	}

	id, err := strconv.Atoi(idBus)
	if (err != nil) {
		return c.JSON(400, map[string]interface{}{
			"message": "Invalid request payload",
		})
	
	}

	business := models.Business{}
	errFind := db.Where(&models.Business{ID: (id)}).First(&business).Error
	if errFind != nil {
		return c.JSON(500, map[string]interface{}{
			"message": "Could not find business",
		})
	
	}
	if business.ID == 0 {
		return c.JSON(404, map[string]interface{}{
			"message": "Business does not exist",
		})
	}


	payload:= new(requests.CreateProduct)
	if err := c.Bind(payload); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	product := models.Products{
		Name :        payload.Name,
		Description : payload.Description,
		Price :       payload.Price,
		BusinessID :  id, 
		Quantity :    payload.Quantity,
		Category :    payload.Category,
		Image :       payload.Image,
		Disabled:    payload.Disabled,
		
	}

	insertErr  := db.Create(&product).Error
	if insertErr != nil {
		return c.JSON(500, map[string]interface{}{
			"message": "Could not create product",
		})
	
	}
	res := make(map[string]interface{})
	res["message"] = "Product created successfully"
	res["product"] = product
	return c.JSON(200, res)




}

func DeleteProduct(c echo.Context, db *gorm.DB) error {
	return nil
}

func UpdateProduct(c echo.Context, db *gorm.DB) error {
	return nil
}

func GetProduct(c echo.Context, db *gorm.DB) error {
	return nil
}

func GetProducts(c echo.Context, db *gorm.DB) error {
	return nil
}


//test functions 
func GetAll (c echo.Context , db *gorm.DB) error {
	data := []models.Products{}
	err := db.Find(&data).Error
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"message": "Could not fetch data",
		})
	}

	return c.JSON(200, data)
}