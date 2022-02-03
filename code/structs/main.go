package main

import "fmt"

// struct Allows defining a data structure
type contactInfo struct {
	// No "," or whatever to separate the different properties
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact contactInfo // Embedding struct indicating the name of the property
	contactInfo // Embedding struct, inferring the name of the property === struct
}

func main() {
	// Declare an object based on a struct is similar to other program language, such as JS, Python, Ruby
	// Several ways to declare
	// 1) Based on the position of the arguments
	contact := contactInfo{"jim@gmail.com", 94000}
	fmt.Println("ContactInfo: ", contact)

	// 2) Indicate all properties
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, // All single line declaring in this way should end with ","
		}, // Even closing an embedded part should end with ","
	}

	// Declaring without initializing
	var alfredo person
	fmt.Println("Not initialized ", alfredo)
	// Printf formats according to a format specifier
	// %+v Deep into the element's properties to return their value
	fmt.Printf("With Printf %+v \n", alfredo)

	// Update property's value
	alfredo.firstName = "Alfredo"
	fmt.Printf("Updated the firstName %+v", alfredo)

	alfredo.updateNameWithoutPointer("Carlos")
	fmt.Println("UpdateNameWithoutPointer: ", alfredo)

	jim.updateName("jimmy")
	jim.print()
}

// Receiver function but without pointer
func (p person) updateNameWithoutPointer(newFirstName string) {
	p.firstName = newFirstName
}

// Function to update the value of the property firstName
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Receiver function
func (p person) print() {
	fmt.Printf("%+v", p)
}
