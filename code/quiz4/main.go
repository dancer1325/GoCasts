package main

import "fmt"

func main() {
	// Array
	arrayOfNumbers := [2]int{1, 2}
	arrayString := [2]string{"Ace of Diamonds", "Five of Shades"}
	fmt.Println(arrayOfNumbers)
	fmt.Println(arrayString)

	// "println" function from builtin package doesn't work for array / slices
	//println(arrayString)

	// Slices

	//Initialize a slice of strings
	cards := []string{"Ace of Diamonds", "Five of Shades"}
	fmt.Println(cards)

	// Slices can be modified their elements
	cardsTwo := append(cards, "Two of Spades") // append doesn't modify the entry elements
	fmt.Println(cards)
	fmt.Println(cardsTwo)

	// Loops
	// Range allows iterating over an array / slices, providing either
	//1) current index
	//2) current element
	for i, card := range cardsTwo {
		fmt.Println("Range with 2 variables", i, card)
	}

	// Range
	// If you don't indicate 2 variables returning the range --> By default it's just the index
	for index := range cardsTwo {
		fmt.Println("Range with just 1 argument", index)
	}

	// Range
	// If we only want the current element --> Use the blank identifier "_"
	for _, element := range cardsTwo {
		fmt.Println("Range with just blank identifier", element)
	}
}
