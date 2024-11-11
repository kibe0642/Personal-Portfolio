package service

import (
	"kibet-kevin-portfolio/internal/model"
	"kibet-kevin-portfolio/internal/repository"
	"log"
)

// SaveContactForm handles the business logic of saving the contact form to the DB
func SaveContactForm(contact model.Contact) error {
	err := repository.SaveContactToDB(contact)
	if err != nil {
		log.Printf("Error saving contact form: %v", err)
		return err
	}
	return nil
}
