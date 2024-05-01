package requests

type CreateProduct struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	Disabled    bool    `json:"disabled"`
}