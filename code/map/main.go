package main

import "fmt"

func main() {

	// Ways to declare a Map
	// 1) Normal type declaration
	var emptyMap map[string]string
	// If you declare a map --> Empty values for each key / value
	printMap(emptyMap)
	// 2) Using make to initialize empty
	emptyMapTwo := make(map[string]string)
	// If you declare a map --> Empty values for each key / value
	printMap(emptyMapTwo)

	// Map to relate the color of the string --> Corresponding RGB value
	colors := map[string]string{
		"red":   "#ff0000", // Different elements of the map are splitted by ","
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["yellow"] = "laksjdf"

	// Delete an entry of the map
	delete(colors, "yellow")

	printMap(colors)
}

func printMap(c map[string]string) {
	//Print all the object
	fmt.Println("Without navigating into each element of the map ", c)

	// Print each element into the map
	// If the map is just declared --> range mapJustDeclared won't fail
	for color, hex := range c {
		// First element   Key of the map
		// Second element   Value of the map
		fmt.Println("Hex code for", color, "is", hex)
	}
}
