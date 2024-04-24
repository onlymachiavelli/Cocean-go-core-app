package handlers

import (
	"fmt"
	"net/http"

	"cocean.com/src/models"
	"cocean.com/src/requests"
	"cocean.com/src/responses"
	"cocean.com/src/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateAdmin(c echo.Context, db *gorm.DB) error {

	userData := new(requests.CreateAdminRequest)
	if err := c.Bind(userData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if userData.Email == "" || userData.FirstName == "" || userData.LastName == "" || userData.Password == "" || userData.Phone == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "All Fields Are Required")
	}

	var admin models.Admins
	if err := db.Where("email = ?", userData.Email).First(&admin).Error; err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Email Already Exists")
	}
	if err := db.Where("phone = ?", userData.Phone).First(&admin).Error; err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Phone Already Exists")
	}

	hashedPass, errorPass := utils.HashPassword(userData.Password)
	if errorPass != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorPass)
	}
	row := &models.Admins{
		Email:     userData.Email,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Phone:     userData.Phone,
		Password:  hashedPass,
	}
	resp := db.Create(row)
	if resp.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, resp.Error)
	}
	created := &responses.CreatedResponse{
		Message: "Admin Has Been Created",
	}
	return c.JSON(http.StatusOK, created)

}


func Login (c echo.Context, db *gorm.DB) error {



	fmt.Println("Hit the login handler")
	payload := new (requests.AdminLoginRequest)
	if err := c.Bind(payload) ; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
		var admin models.Admins
		if err := db.Where("email = ?", payload.Email).First(&admin).Error; err != nil {
			responseErr := responses.ErrorResponse{
				Message: "Admin user Not Found",
			}
			return echo.NewHTTPError(http.StatusNotFound, &responseErr)
		}
		if (admin.ID == 0) {
			errResponse := responses.ErrorResponse{
				Message: "Admin user Not Found",
			}
			return echo.NewHTTPError(http.StatusNotFound, &errResponse)
		}
		if !utils.Verify(payload.Password, admin.Password) {
			errResponse := responses.ErrorResponse{
				Message: "Invalid Email Password",
			}
			return echo.NewHTTPError(http.StatusUnauthorized, &errResponse)
		}
		token, err := utils.GenerateToken(int(admin.ID))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]string{
			"jwt": token,
		}) 
	
}