package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"io/fs"
	"os"

	"github.com/vezzalinistefano/learning-htmx/models"
)

type contactRepository struct {
	contacts []models.Contact
}

var ContactsRepository contactRepository

const pageSize = 50

func init() {
	jsonData, err := os.ReadFile("./data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(jsonData, &ContactsRepository.contacts)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (c *contactRepository) storeToJson() error {
	jsonBytes, err := json.Marshal(c.contacts)
	if err != nil {
		fmt.Println("Unable to write to json")
		return errors.New("Unable to store to JSON file")
	}

	fmt.Println("Saved to JSON")
	_ = os.WriteFile("./data.json", jsonBytes, fs.FileMode(0644))
	return nil

}

// Private Methods

func (c *contactRepository) getIndexById(id int) (*int, error) {
	for idx, contact := range c.contacts {
		if contact.Id == id {
			return &idx, nil
		}
	}
	return nil, errors.New("Contact not found!")
}

// Public Methods

func (c *contactRepository) Count() int {
    return len(c.contacts)
}

func (c *contactRepository) Search(query string) []models.Contact {
	var r []models.Contact
	for _, contact := range c.contacts {
		if strings.Contains(contact.First, query) || strings.Contains(contact.Last, query) {
			r = append(r, contact)
		}
	}
	return r
}

func (c *contactRepository) GetAll(page int) []models.Contact {
	start := (page - 1) * pageSize
	end := start + pageSize
	if end >= len(c.contacts) {
		end = len(c.contacts) - 1
	}
	return c.contacts[start:end]
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
	contact.Id = len(c.contacts) + 1
	c.contacts = append(c.contacts, contact)

	err := c.storeToJson()
	if err != nil {
		return
	}
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
