package handlers

import (
	"strconv"

	"cocean.com/src/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetMyClients(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	//convert id to atoi
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, "Invalid request payload")
	}
	//find business 
	business := models.Business{}
	errFind := db.Where(&models.Business{ID: idInt}).First(&business).Error
	if errFind != nil {
		return c.JSON(500, "Could not find business")
	}
	if business.ID == 0 {
		return c.JSON(404, "Business does not exist")
	}
	//find clients
	clients := []models.Clients{}
	errFindClients := db.Where(&models.Clients{Business: idInt}).Find(&clients).Error
	if errFindClients != nil {
		return c.JSON(500, "Could not find clients")
	}
	return c.JSON(200, clients)
}


//delete one 

//send em all email 


//send one email 