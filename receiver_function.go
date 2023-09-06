package main

import "fmt"

// Define the custom type Person based on a struct
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

// Receiver function for the Person type
func (p Person) GetFullName() string {
    return p.FirstName + " " + p.LastName
}

func main() {
    // Create a Person instance
    person := Person{
        FirstName: "Saurabh",
        LastName:  "Joshi",
        Age:       26,
    }

    // Call the receiver function
    fullName := person.GetFullName()

    fmt.Println("Full Name:", fullName)
}
