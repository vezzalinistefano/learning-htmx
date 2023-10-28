package repositories

import (
	"github.com/vezzalinistefano/learning-htmx/models"
)

type contactRepository struct {
	contacts []models.Contact
}

var ContactsRepository contactRepository

func init() {
	ContactsRepository.contacts = append(ContactsRepository.contacts, models.Contact{Id: 1, First: "Alice", Last: "Smith", Phone: "+1-555-555-5555", Email: "alice.smith@example.com"})
	ContactsRepository.contacts = append(ContactsRepository.contacts, models.Contact{Id: 2, First: "Bob", Last: "Jones", Phone: "+1-666-666-6666", Email: "bob.jones@example.com"})
	ContactsRepository.contacts = append(ContactsRepository.contacts, models.Contact{Id: 3, First: "Carol", Last: "Williams", Phone: "+1-777-777-7777", Email: "carol.williams@example.com"})
}

func (c *contactRepository) GetAll() []models.Contact {
	return ContactsRepository.contacts
}
