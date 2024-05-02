package utils

import (
	"fmt"

	"cocean.com/src/models"
	"cocean.com/src/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func init() {
	fmt.Println("Hit the connect to db file")
}
func Connect(config types.DBConfig) (*gorm.DB, error) {
	fmt.Println("The config is : " , config)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Europe/Berlin",
		config.Host, config.User, config.Pass, config.Name, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	errMig := db.AutoMigrate(&models.Admins{})
	if errMig != nil {
		return nil, errMig
	}

	errMigBusiness := db.AutoMigrate(&models.Business{})
	if errMigBusiness != nil {
		return nil, errMigBusiness
	}

	errProds := db.AutoMigrate(&models.Products{})
	if errProds != nil {
		return nil, errProds
	}


	//orders 
	errOrders := db.AutoMigrate(&models.Orders{})
	if errOrders != nil {
		return nil, errOrders
	}

	//clients 

	errClients := db.AutoMigrate(&models.Clients{})
	if errClients != nil {
		return nil, errClients
	}

	return db, nil
	
}