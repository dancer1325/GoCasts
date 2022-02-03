package main

// Since no type to return is indicated --> void
func main() {

	// Declaring variables
	// 1) Indicating explicitly the type
	//var card string = "Ace of Spades"
	// 2) Initializing the variable and inferring the type
	//card :=  "Ace of Spades"
	//fmt.Println(card)

	// For reassigning a value
	// 1) Not allowed using :=
	//card := "Five of Diamonds"
	// 2) Use simply =
	//card = "Five of Diamonds"
	//fmt.Println(card)

	cards := newDeck()
	// Similar in this case, to declare an array of strings
	// cards := []string

	//cards.toString()
	//cards.saveToFile("mycards")
	//newDeckFromFile("blabla")  // Check how it works
	cards.shuffle()
	cards.print()

	// Invoking the function with the wrong arguments --> You will get an error
	// deal("aa",2)
	// Capture several types returned
	hand, remainingDeck := deal(cards, 4)
	hand.print()
	remainingDeck.print()
}
