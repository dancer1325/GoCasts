package main

import (
	"fmt"
	"io/ioutil" // Import just the specific required subpackage
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

//Not possible to declare and initialize variables with := outside func body
// type variableDeclaredOutside := 20
//variableDeclaredOutside := 20
// Possible to declare and initialize variables without using :=
var variableDeclaredOutside = 20

// Create a new Deck of playing cards
// This function wouldn't need to be func receiver since it creates new "deck" type
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// Range without using the index --> Use the blank identifier "_"
	for _, suit := range cardSuits {
		// Range without using the index --> Use the blank identifier "_
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver function
// Firstly appear the arguments of the function --> Any variable defined with that type can access to that method
func (d deck) print() {
	// Range with 2 variables indicated
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal a hand, and splitting the cards in 2 array of cards
// Return 2 type "deck"
// It wasn't indicated as a receiver function because it would be ambiguous
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Receiver function, since it makes sense to have toString method as a method own to deck type
func (d deck) toString() string {
	//[]string(d)   Convert deck to []string
	sliceString := []string(d)
	//fmt.Println("Casting deck to []string: ", sliceString)

	oneString := strings.Join(sliceString, ",")
	//fmt.Println("Joining in a simple string, separated by , : ", oneString)
	return oneString
}

// Receiver function, since it makes sense to have saveToFile method as a method own to deck type
// Return error, because WriteFile returns that type
func (d deck) saveToFile(filename string) error {
	// []byte(...)  accepts string in order to be casted
	// 0666   Anyone can read and write the file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	content, err := ioutil.ReadFile(filename)
	if err != nil { // In case of error in ReadFile --> It returns 'nil' === no value, in the error variable
		// Options about what to do in case some error happen
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1) // Indicates some error happened
	}

	s := strings.Split(string(content), ",")
	return deck(s) // Type conversion. It's possible also for custom types
}

// Receiver function
func (d deck) shuffle() {

	// 1) Same seed generating the values
	// It will always generate the same random numbers
	//for i := range d { // It's possible to iterate over the deck's properties.
	//	newPosition := rand.Intn(len(d) - 1) // Return a non-negative pseudo-random number in the half-open internal [0, len(d))
	//
	//	d[i], d[newPosition] = d[newPosition], d[i]
	//	fmt.Println("index: ", i)
	//	fmt.Println("newPosition: ", newPosition)
	//}

	// 2) Different seed to generate the values
	// It will generate different random numbers
	// time.Now().UnixNano()  Returns an int64
	// Source is a type in GO, which is an interface with properties: A) Int63, B) Seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) // Generate rand. Source is the seed for the generator

	// Range with just 1 element indicated --> The index of the iteration
	for i := range d { // It's possible to iterate over the deck's properties.
		newPosition := r.Intn(len(d) - 1) // Return a non-negative pseudo-random number in the half-open internal [0, len(d))

		d[i], d[newPosition] = d[newPosition], d[i]
		fmt.Println("index: ", i)
		fmt.Println("newPosition: ", newPosition)
	}
}
