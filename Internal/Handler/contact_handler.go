package handler

import (
	"net/http"
	"kibet-kevin-portfolio/internal/controller"
)

// ContactHandler handles contact form submissions
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle form submission
		controller.HandleContactForm(w, r)
	} else {
		// Render the contact page
		http.ServeFile(w, r, "templates/contact.html")
	}
}
