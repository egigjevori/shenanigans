package main

import "fmt"

func main() {
    phonebook := NewTree[string]()

    phonebook.Put("Alice", "123-456-7890")
    phonebook.Put("Bob", "234-567-8901")
    phonebook.Put("Charlie", "345-678-9012")
    phonebook.Put("David", "456-789-0123")
    phonebook.Put("Diana", "567-890-1234")

    contactName := "Alice"
    if phoneNumber, exists := phonebook.Search(contactName); exists {
        fmt.Printf("Phone number for %s: %s\n", contactName, phoneNumber)
    } else {
        fmt.Printf("Contact %s not found in the phonebook\n", contactName)
    }

    prefix := "Da"
    if phonebook.StartsWith(prefix) {
        fmt.Printf("Contacts starting with '%s':\n", prefix)
        contacts := phonebook.GetWordsWithPrefix(prefix)
        for _, entry := range contacts {
            fmt.Printf("Name: %s, Phone Number: %s\n", entry.Word, entry.Value)
        }
    } else {
        fmt.Printf("No contacts start with the prefix '%s'\n", prefix)
    }

    deleteContact := "Charlie"
    phonebook.Delete(deleteContact)
    if _, exists := phonebook.Search(deleteContact); !exists {
        fmt.Printf("Contact %s successfully deleted from the phonebook\n", deleteContact)
    } else {
        fmt.Printf("Failed to delete contact %s from the phonebook\n", deleteContact)
    }
}
