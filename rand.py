import json
import random


def generate_fake_contact(i):
    first_names = ["Alice", "Bob", "Carol", "Dave", "Eve",
                   "Frank", "George", "Helen", "Irene", "John", "Karen"]
    last_names = ["Smith", "Jones", "Williams", "Brown", "Johnson",
                  "Miller", "Taylor", "Davis", "Anderson", "Wilson", "Thomas"]
    phone_numbers = ["+1-555-555-5555", "+1-555-555-5556", "+1-555-555-5557",
                     "+1-555-555-5558", "+1-555-555-5559", "+1-555-555-5560"]
    email_domains = ["example.com", "gmail.com",
                     "yahoo.com", "hotmail.com", "outlook.com"]

    first_name = random.choice(first_names)
    last_name = random.choice(last_names)
    phone_number = random.choice(phone_numbers)
    email_domain = random.choice(email_domains)
    email = first_name + "." + last_name + "@" + email_domain

    return {
        "id": i+1,
        "first": first_name,
        "last": last_name,
        "phone": phone_number,
        "email": email,
        "Errors": {}
    }


# Generate 50 fake contacts
fake_contacts = []
for i in range(50):
    fake_contacts.append(generate_fake_contact(i))

# Save the fake contacts to a JSON file
with open("data.json", "w") as f:
    json.dump(fake_contacts, f)
