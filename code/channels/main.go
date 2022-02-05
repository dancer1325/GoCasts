package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Using the package http --> Require declare all the urls with http included
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Ways to execute code
	// 1) Sequentially
	//for _, link := range links {
	//	checkLinkWithoutChannels(link)
	//}

	// 2) Creating other Go routines, without using channels --> Concurrent or parallelly (depends on the number of CPU core
	// It won't work as it's expected, since the main Go routine can close previously to the end of the child Go routines
	//for _, link := range links {
	//	go checkLinkWithoutChannels(link)
	//}

	// 3) Creating other Go routines and using channels --> Concurrent or parallelly (depends on the number of CPU core)
	// Initialize a channel of type string. Since we have declared here --> it's scope is into the main function
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // If you want that the go routine has got a channel --> Add as argument
	}

	// Wait for Hearing if some value is sent in the channel to print out. Blocking part, waiting for receiving messages from a channel
	// But, once some of the child go routines send data to the channel --> With the next line just the main go routine would finish, BUT meantine
	// the other child go routines would stay forever.
	// It would hear just for 1! channel
	//fmt.Println(<- c)

	// If we create 2 --> It would try to hear 2 different channels ones
	//fmt.Println(<- c)
	//fmt.Println(<- c)

	// Receive and hear the specific number of channels to which one, information has been sent
	// It continues being blocking code, since not all channels are listened to at the same time. They are heard one to one
	//for i := 0; i < len(links); i++ {
	//	fmt.Println(<-c)
	//}

	// Check each link a second time, once it's heard in the channel
	// This loop doesn't continue working to the eternity, just 5 times is heard some data send to the channel
	//for i := 0; i < len(links); i++ {
	//	go checkLink(<-c, c)
	//}

	// Infinite loop checking continuously
	//for {
	//	go checkLink(<-c, c)
	//}

	// Equivalent to the previous code
	// range c   Wait for to receive some element sent in the channel
	// l         Each element received in the channel
	//for l := range c {
	//	go checkLink(l, c)
	//}

	for l := range c { // Each time that a new value is heard, we are overriding the value l in the same memory address
		// 1) Pausing the main go routine
		//time.Sleep(5 * time.Second) // It would pause the current go routine, so the main go routine --> No other message in other channel can be heard
		//go checkLink(l, c)

		// 2) Literal function created in another go routine, but without specifying argument in the anonymous function
		// It doesn't work properly because variables in different go routines shouldn't be referred, and reference type is being used
		//go func() {
		//	time.Sleep(5 * time.Second) // Pause the current go routine
		//	checkLink(l, c)
		//}(l) // ()  --> Invoke this literal function

		// 3) Literal function created in another go routine, passing a specific argument to the function
		// It works as it's expected, since it's passed a copy value --> Although the memory address changes in the main go routine, in the child go routine is used the copy value originally received
		go func(link string) {
			time.Sleep(5 * time.Second) // Pause the current go routine
			checkLink(link, c)
		}(l) // (l)  --> Invoke this literal function
	}
}

// Function to check the links sequentially
//func checkLinkWithoutChannels(link string) {
//	// Response object isn't important for us, since we only want to check if it's possible to reach or not the url
//	_, err := http.Get(link)
//	if err != nil {
//		fmt.Println(link, "might be down!")
//		return
//	}
//
//	fmt.Println(link, "is up!")
//}

func checkLink(link string, c chan string) {
	// Response object isn't important for us, since we only want to check if it's possible to reach or not the url

	// Overcome the problem in which part to place the pause
	//time.Sleep(5 * time.Second) // Not good idea to pause in the checkLink function, although it's in the child go routine, because it's nor their current goal

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // Send link to the channel. To identify in the main go routine, the link already checked
		return
	}

	fmt.Println(link, "is up!")
	c <- link // Send link to the channel. To identify in the main go routine, the link already checked
}
