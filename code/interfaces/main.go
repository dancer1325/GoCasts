package main

import "fmt"

type bot interface {
	getGreeting() string
}

// Struct types, although no properties to them, used because we want to create some functions related to it
// Interfaces are implicit === You don't specify by code explicitly that a type extends for the interface
type englishBot struct{}
type spanishBot struct{}

func main() {
	// Declaring a value of a struct which hasn't got any property
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// Implementatin of the interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Although the argument are different of the next functions, GO compiler complains (Different to Java)
//func printGreeting(en englishBot) {
//	fmt.Println(en.getGreeting())
//}
//
//func printGreeting(es spanishBot) {
//	fmt.Println(es.getGreeting())
//}

//func (eb englishBot) getGreeting() string {
// Since it's not used at all receiver argument --> Remove it and leave the type just
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

// Since the type englishBot has got a function named "getGreeting" with the same arguments and return type -->
// englishBot is a bot --> englishBot can be passed to printGreeting()

//func (es spanishBot) getGreeting() string {
// Since it's not used at all receiver argument --> Remove it and leave the type just
func (spanishBot) getGreeting() string {
	return "Hola!"
}

// Since the type spanishBot has got a function named "getGreeting" with the same arguments and return type -->
// spanishBot is a bot --> spanishBot can be passed to printGreeting()
