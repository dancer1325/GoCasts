package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Args
	// Variable to hold the CLI arguments
	// os.Args[0]  Program name
	fmt.Println("Program name: ", os.Args[0])

	// os.Args[1]  Return the CLI argument. If no argument is indicated --> It will throw an error
	if len(os.Args) == 1 {
		fmt.Println("No argument added to the console about the name of the file to read")
		os.Exit(1)
	}

	fmt.Println("CLI argument: ", os.Args[1])

	// f  *File. Value of the type File
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Copy the *File to the os.Stdout
	io.Copy(os.Stdout, f)
}
