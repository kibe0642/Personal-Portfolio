package controller

import (
	"kibet-kevin-portfolio/internal/model"
	"kibet-kevin-portfolio/internal/service"
	"net/http"
)

// HandleContactForm processes the form submission and sends it to the service
func HandleContactForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	contact := model.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	// Call service to process the form
	if err := service.SaveContactForm(contact); err != nil {
		http.Error(w, "Failed to save contact form", http.StatusInternalServerError)
		return
	}

	// Success response
	w.Write([]byte("Thanks for reaching out! We'll get back to you soon."))
}
