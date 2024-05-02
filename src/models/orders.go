package models

type Orders struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Business int    `json:"business"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Comment  string `json:"comment,omitempty"`

	Total    int    `json:"total"`
	Products string `json:"products"`

	Status    string `json:"status"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime"`
}