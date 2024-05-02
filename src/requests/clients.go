package requests

type SendMail struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}