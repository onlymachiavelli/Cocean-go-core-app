package requests

type CreateOrder struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Comment  string `json:"comment,omitempty"`
	Business int    `json:"business"`
	Total    int    `json:"total"`
	Products string `json:"products"`
	Status   string `json:"status"`
}