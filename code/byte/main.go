package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a1 byte = 97
	var a2 byte = 98
	var a3 byte = 99

	fmt.Println("Print the byte directly ")
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println("Print the character representation of the byte ")
	// %c  Format verb -- Print the character representation of the byte
	fmt.Printf("%c\n", a1)
	fmt.Printf("%c\n", a2)
	fmt.Printf("%c\n", a3)

	// Without specifying the byte type explicitly
	var a4 byte = 97
	var a5 = 98
	a6 := 'c'

	fmt.Println("Print the byte directly ")
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)

	fmt.Println("Print the character representation of the byte ")
	// %c  Format verb -- Print the character representation of the byte
	fmt.Printf("%c\n", a4)
	fmt.Printf("%c\n", a5)
	fmt.Printf("%c\n", a6)

	fmt.Println("Print the type of each variable ")
	// TODO: Which formatting to use to extract exactly the type of the variable?
	fmt.Printf("Type of a4= ", reflect.TypeOf(a4)) // If you declare byte --> uint8
	fmt.Printf("Type of a5= ", reflect.TypeOf(a5)) // If you declare and initialize with var --> int
	fmt.Printf("Type of a6= ", reflect.TypeOf(a6)) // If you declare and initialize with ":=" --> int32

	// String --> bytes
	fmt.Println("\n")
	fmt.Println([]byte("Name")) // Type conversion to []byte

	// Bytes --> String
	dataByte := []byte{102, 97, 108, 99, 111, 110}

	fmt.Println(dataByte)
	fmt.Println(string(dataByte)) // Type conversion to string
}
