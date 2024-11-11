package repository

import (
	"kibet-kevin-portfolio/internal/model"
	"kibet-kevin-portfolio/db"
	"log"
)

// SaveContactToDB saves the contact form data to the MSSQL database
func SaveContactToDB(contact model.Contact) error {
	db := db.GetDB()

	query := `INSERT INTO Contacts (Name, Email, Message) VALUES (?, ?, ?)`
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Message)
	if err != nil {
		log.Printf("Error inserting contact data into DB: %v", err)
		return err
	}
	return nil
}
