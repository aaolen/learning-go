// Write a function called prefixer that has an input parameter of type string
// and returns a function that has an input parameter of type string and
// returns a string. The returned function should prefix its input with the
// string passed into prefixer. Use the following main function to test
// prefixer:
// func main() {
//     helloPrefix := prefixer("Hello")
//     fmt.Println(helloPrefix("Bob")) // should print Hello Bob
//     fmt.Println(helloPrefix("Maria")) // should print Hello Maria
// }

package main

import (
	"fmt"
)

func prefixer(prefix string) func(string) string {
	return func(input string) string {
		return fmt.Sprintf("%s %s", prefix, input)
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
