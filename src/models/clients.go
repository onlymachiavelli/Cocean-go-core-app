package models

type Clients struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Totalorders int    `json:"totalorders"`
	Business    int    `json:"business"`
	CreatedAt   int    `json:"created_at" gorm:"autoCreateTime"`
}