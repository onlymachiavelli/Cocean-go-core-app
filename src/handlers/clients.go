package handlers

import (
	"strconv"

	"cocean.com/src/models"
	"cocean.com/src/requests"
	"cocean.com/src/utils"
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






func Broadcast (c echo.Context, db*gorm.DB) error{

	idBus := c.Param("id")
	
	


	id, err := strconv.Atoi(idBus)
	if err != nil {
		return c.JSON(400, "Invalid request payload")
	}

	business := models.Business{}
	errFind := db.Where(&models.Business{ID: id}).First(&business).Error
	if errFind != nil {
		return c.JSON(500, "Could not find business")
	}


	if business.ID == 0 {
		return c.JSON(404, "Business does not exist")
	}


	clients := []models.Clients{}
	errFindClients := db.Where(&models.Clients{Business: id}).Find(&clients).Error
	if errFindClients != nil {
		return c.JSON(500, "Could not find clients")
	}


	payload := new(requests.SendMail)

	if err := c.Bind(payload); err != nil {
		return c.JSON(400, "Invalid request payload")
	}


	//send email to all clients
	for _, client := range clients {
		//send email to client
		//send email to client
		errSend := utils.SendMail(client.Email, payload.Subject,payload.Message)
		if errSend != nil {

			return c.JSON(500, "Could not send email")
		}
	}

	

	return c.JSON(200, "Email sent to all clients")




}