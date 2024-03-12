// Write two functions. The UpdateSlice function takes in a []string and a
// string. It sets the last position in the passed-in slice to the passed-in
// string. At the end of UpdateSlice, print the slice after making the change.
// The GrowSlice function also takes in a []string and a string. It appends
// the string onto the slice. At the end of GrowSlice, print the slice after
// making the change. Call these functions from main. Print out the slice
// before each function is called and after each function is called. Do you
// understand why some changes are visible in main and why some changes are
// not?

package main

import "fmt"

func UpdateSlice(sSlice []string, s string) {
	sSlice[len(sSlice)-1] = s
	fmt.Println("Result", sSlice)
}

func GrowSlice(sSlice []string, s string) {
	newSlice := append(sSlice, s)
	fmt.Println("Result: ", newSlice)
}

func main() {
	slice := []string{"first", "second", "third"}

	fmt.Println("Before update: ", slice)
	UpdateSlice(slice, "update")
	fmt.Println("After update: ", slice)

	fmt.Println()

	fmt.Println("Before grow: ", slice)
	GrowSlice(slice, "grow")
	fmt.Println("After grow: ", slice)
}
