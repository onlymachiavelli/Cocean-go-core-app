package models

type Business struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"not null"`
	Address     string `json:"address" gorm:"not null"`
	Phone       string `json:"phone" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	Owner       int    `json:"owner" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Category    string `json:"category" gorm:"not null"`
	CreatedAt   int    `json:"created_at" gorm:"autoCreateTime"`
}