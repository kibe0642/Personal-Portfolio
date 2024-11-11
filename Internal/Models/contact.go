package model

// Contact represents the structure of the contact form data
type Contact struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
