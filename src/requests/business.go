package requests

import "fmt"

func init() {
	fmt.Println("hit the business reuqests package")
}



//create business only 

type CreateBusiness struct {
	Name        string `json:"name" gorm:"not null"`
	Address     string `json:"address" gorm:"not null"`
	Phone       string `json:"phone" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Category    string `json:"category" gorm:"not null"`
	Photo       string `json:"photo" gorm:"not null"`
	Owner 	 int    `json:"owner" gorm:"not null"`

}