package utils

import (
	"log"
	"os"

	"github.com/go-gomail/gomail"
	"github.com/joho/godotenv"
)

func SendMail(target string , subject string, message string ) error{
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	google := os.Getenv("GOOGLE")
	pass := os.Getenv("GMAILPASS")
	
	htmlTemplate := `
	<!DOCTYPE html>
	<html> 
	  <head>
		<title>Email Template</title>
		<style>
		  /* Add your CSS styles here */
		  body {
			font-family: Arial, sans-serif;
			background-color: #f4f4f4;
			padding: 20px;
		  }
		  .container {
			max-width: 600px;
			margin: 0 auto;
			background-color: #fff;
			padding: 30px;
			border-radius: 8px;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		  }
		  h1 {
			color: #333;
		  }
		  p {
			color: #666;
		  }
		</style>
	  </head>
	  <body>
		<div class="container">
		  <h1>` + message+ `</h1>
		</div>
	  </body>
	</html>
	`
	
	
	m := gomail.NewMessage()
	m.SetHeader("From", google)
	m.SetHeader("To", target)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", htmlTemplate)
	d := gomail.NewDialer("smtp.gmail.com", 587, google, pass)
	return d.DialAndSend(m)

}



