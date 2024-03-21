// Write a program that defines a struct called Employee with three fields:
// firstName, lastName, and id. The first two fields are of type string, and
// the last field (id) is of type int. Create three instances of this struct
// using whatever values you’d like. Initialize the first one using the struct
// literal style without names, the second using the struct literal style with
// names, and the third with a var declaration. Use dot notation to populate
// the fields in the third struct. Print out all three structs.

package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	a := Employee{
		"Johnny",
		"Appleseed",
		0,
	}
	b := Employee{
		"Johnny",
		"Appleseed",
		1,
	}
	var c Employee
	c.firstName = "Johnny"
	c.lastName = "Appleseed"
	c.id = 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
