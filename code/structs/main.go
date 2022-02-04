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

	fmt.Println("Jim. Previous to update anthing ", jim)

	// Ways to avoid passing by value in go when you invoke a function with having the pointer of the type
	// 1) Simply by the type
	jim.updateName("jimmy")
	jim.print()

	// 2) Define the pointer explicitly
	var jimPointer *person
	// Access to the memory address of the pointer
	jimPointer = &jim
	// TODO: Although you are using a pointer --> Is there no way to print the memory address?
	fmt.Println(" jimPointer: ", jimPointer)
	jimPointer.print()
}

// Receiver function but without pointer --> Afterwards the element's property hasn't been updated
// Go is a pass value language -->
// 1) Although you are doing jim.updateNameWithoutPointer --> It's creating another person in the memory
// 2) Update the information on the new person created
func (p person) updateNameWithoutPointer(newFirstName string) {
	p.firstName = newFirstName
}

// Function to update the value of the property firstName, but for a specific memory address, not for the copy pass value
// * as argument, because it's the pointer of the typer person
func (pointerToPerson *person) updateName(newFirstName string) {
	// TODO: Is necessary to add "*" to the variable pointerToPerson? I don't see any difference in the result
	//(*pointerToPerson).firstName = newFirstName
	(pointerToPerson).firstName = newFirstName
}

// Receiver function
func (p person) print() {
	fmt.Printf("%+v", p)
}
