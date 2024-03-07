// 1. Write a for loop that puts 100 random numbers between 0 and 100 into an int
// slice.
// 2. Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
// If the value is divisible by 2, print “Two!”
// If the value is divisible by 3, print “Three!”
// If the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
// Otherwise, print “Never mind”.

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// 1:
	intSlice := make([]int, 100)
	for i := range intSlice {
		intSlice[i] = rand.IntN(100)
	}
	fmt.Println(intSlice)

	// 2:
	for _, v := range intSlice {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Print("Six! ")
		case v%2 == 0:
			fmt.Print("Two! ")
		case v%3 == 0:
			fmt.Print("Three! ")
		default:
			fmt.Print("Nevermind. ")
		}
	}
	fmt.Println()

}
