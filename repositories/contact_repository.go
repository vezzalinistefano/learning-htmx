package repositories

import (
	"errors"
	"strings"

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

// Private Methods

func (c *contactRepository) search(query string) []models.Contact {
	var r []models.Contact
	for _, contact := range c.contacts {
		if strings.Contains(contact.First, query) || strings.Contains(contact.Last, query) {
			r = append(r, contact)
		}
	}
	return r
}

func (c *contactRepository) getIndexById(id int) (*int, error) {
	for idx, contact := range c.contacts {
		if contact.Id == id {
			return &idx, nil
		}
	}
	return nil, errors.New("Contact not found!")
}

// Public Methods

func (c *contactRepository) GetAll(query string) []models.Contact {
	if query == "" {
		return c.contacts
	} else {
		return c.search(query)
	}
}

func (c *contactRepository) GetByContactID(id int) (*models.Contact, error) {
	for _, contact := range c.contacts {
		if contact.Id == id {
			return &contact, nil
		}
	}
	return nil, errors.New("Contact not found!")
}

func (c *contactRepository) InsertContact(contact models.Contact) {
	c.contacts = append(c.contacts, contact)
}

func (c *contactRepository) EditContact(contact models.Contact) {
	if idx, err := c.getIndexById(contact.Id); err == nil {
		c.contacts[*idx] = contact
	} else {
		return
	}
}

func (c *contactRepository) DeleteContactById(id int) {
	if idx, err := c.getIndexById(id); err == nil {
		c.contacts = append(c.contacts[:*idx], c.contacts[*idx+1:]...)
	} else {
		return
	}
}
