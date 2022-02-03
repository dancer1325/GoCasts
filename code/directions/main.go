package main

import "fmt"

func main() {
	var intSlice []int

	// Go loop patterns https://yourbasic.org/golang/for-loop/
	// 1) Three-component loop
	//for i := 0; i<=10; i++ {
	//	intSlice = append(intSlice, i)
	//}
	//fmt.Println(intSlice)
	// 2) While loop
	index := 0
	for index <= 10 {
		intSlice = append(intSlice, index)
		index++
	}
	//fmt.Println(intSlice)

	for _, element := range intSlice {
		if element%2 == 0 {
			fmt.Println(element, " is even")
		} else {
			fmt.Println(element, " is odd")
		}
	}
}
