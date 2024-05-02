package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"cocean.com/src/models"
	"cocean.com/src/requests"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateOrder(c echo.Context, db *gorm.DB) error {

	businessId := c.Param("id")
	//convert it to int 
	
	id ,errId := strconv.Atoi(businessId)
	
	if errId != nil {
		return c.JSON(400, "Missing parameter: bus_id")
	}
	if id == 0 {
		return c.JSON(400, "Invalid request payload")
	
	}

	payload := new(requests.CreateOrder)
	if err := c.Bind(payload); err != nil {
		return c.JSON(400, "Invalid request payload")
	}


	//find the business 
	business := models.Business{}
	errFind := db.Where(&models.Business{ID: id}).First(&business).Error
	if errFind != nil {
		return c.JSON(500, "Could not find business")
	}
	if business.ID == 0 {
		return c.JSON(404, "Business does not exist")
	}

	order := &models.Orders{
		Business: id,
		Name: payload.Name,
		Phone: payload.Phone,
		Email: payload.Email,
		Address: payload.Address,
		Comment:  payload.Comment, 
		Total: payload.Total ,
		Products: payload.Products,
		Status: "pending",
	}


	//collect user info 

	err := db.Create(order).Error
	if err != nil {
		return c.JSON(500, "Could not create order")
	}

	//find the user 
	client := models.Clients{}
	//find by business and email 
	errFind = db.Where(&models.Clients{Business: id, Email: payload.Email}).First(&client).Error
	
	if errFind != nil {
		fmt.Println("that's whatsup")
	}
	if client.ID == 0 {
		//create the user 
		client = models.Clients{
			
			Name : payload.Name,
			Phone : payload.Phone,
			Email : payload.Email,
			Address : payload.Address,
			Totalorders : 1,
			Business : id,
		}
		//save the user 
		err = db.Create(&client).Error
		if err != nil {
			return c.JSON(500, "Could not create user")
		}
		
	}
		//update the user 
		client.Totalorders = client.Totalorders + 1
		err = db.Save(&client).Error
		if err != nil {
			return c.JSON(500, "Could not update user")
		}
		//save the order
		err = db.Create(&order).Error
		if err != nil {
			fmt.Println("saved but damn!!")
		}
		
		//response 
		res := make(map[string]interface{})
		res["message"] = "Order created successfully"
		res["order"] = order
		return c.JSON(200, res)

	

}


func GetOrders(c echo.Context, db *gorm.DB) error {

	businessId := c.Param("id")
	//convert it to int
	id, errId := strconv.Atoi(businessId)
	if errId != nil {
		return c.JSON(400, "Missing parameter: bus_id")
	}

	if id == 0 {
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

	//find orders by business
	orders := []models.Orders{}
	errFindOrders := db.Model(&models.Orders{}).Where("business = ?", id).Find(&orders).Error
	if errFindOrders != nil {
		return c.JSON(500, "Could not find orders")
	}

	type P struct {
		Prod models.Products `json:"products"`
		Quantity int `json:"quantity"`
	}
	type OrderWithProducts struct {
		models.Orders
		
		Products []P `json:"products"`

	}

	ordersWithProducts := []OrderWithProducts{}
	type jsonTst struct{
		ID int `json:"id"`
		Quantity int `json:"quantity"`
	}
	
	for _, order := range orders {
		var ids []jsonTst
		err := json.Unmarshal([]byte(order.Products), &ids)
		fmt.Println(err)
		if err != nil {
			return c.JSON(500, "Could not parse products")
		}
		

		products := []P{}
		for _, id := range ids {
			product := models.Products{}
			errFindProduct := db.Where(&models.Products{ID: id.ID}).First(&product).Error
			if errFindProduct != nil {
				return c.JSON(500, "Could not find product")
			}
			products = append(products, P{Prod: product, Quantity: id.Quantity})
		}

		ordersWithProducts = append(ordersWithProducts, OrderWithProducts{Orders: order, Products: products})
	}

	
	for _, order := range ordersWithProducts {
		for _, product := range order.Products {
			product.Prod.Quantity = product.Prod.Quantity - product.Quantity
			err := db.Save(&product.Prod).Error
			if err != nil {
				return c.JSON(500, "Could not update stock")
			}
		}
	}

	return c.JSON(200, ordersWithProducts)
}


func GetOrder(c echo.Context , db *gorm.DB) error {
	return nil 
}

func UpdateOrder (c echo.Context , db *gorm.DB) error {
	return nil 
}


func DeleteOrder (c echo.Context , db *gorm.DB) error {
	return nil 
}

