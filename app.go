package main

import (
	"fmt"
	"os"

	route "cocean.com/src/routes"
	"cocean.com/src/types"
	"cocean.com/src/utils"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// func injectCors(e *echo.Echo) {
// 	devMode := true
// 	if (devMode) {
// 		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 			AllowOrigins: []string{"*"},
// 			AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
// 		  }))

// 	} else {
// 		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 			AllowOrigins: []string{"*"},
// 			AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
// 		  }))
// 	}
// }

func init() {
	fmt.Println("Hit the main file")
}



func main () {



	//cors 
	
	loadEnvErr := godotenv.Load()
	if (loadEnvErr != nil) {
		panic(loadEnvErr)
	}
	configuration := types.DBConfig{
		 Port: os.Getenv("DBPORT"),
		 Name: os.Getenv("DBNAME"),
		 Pass: os.Getenv("PASS"),
		 Host: os.Getenv("HOST"),
		 User: os.Getenv("USER"),

	}

	db, err := utils.Connect(configuration)
	if err != nil {
		panic(err)
	}

		fmt.Println("The db is : " , db)

	
	e := echo.New()
	// injectCors(e)



	
	if (db!= nil) {
		fmt.Println("Connected to the database")

		//health groupe 
		healthGroup := e.Group("/health")
		route.HealthRoute(healthGroup)

		//admin groupe
		adminGroup := e.Group("/me")
		route.AdminRoutes(adminGroup, db)
		


	}
	e.Logger.Fatal(e.Start(":8000"))

}