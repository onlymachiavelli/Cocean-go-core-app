package utils

import (
	"fmt"

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

	

	return db, nil
}