package models

type Products struct {
	ID          int     `json:"id"`
	BusinessID  int     `json:"business_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	Disabled    bool    `json:"disabled"`
	CreatedAt   int     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int     `json:"updated_at" gorm:"autoCreateTime"`
}