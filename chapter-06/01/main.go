// Create a struct named Person with three fields: FirstName and LastName of
// type string and Age of type int. Write a function called MakePerson that
// takes in firstName, lastName, and age and returns a Person. Write a second
// function MakePersonPointer that takes in firstName, lastName, and age and
// returns a *Person. Call both from main. Compile your program with go build
// -gcflags="-m". This both compiles your code and prints out which values
// escape to the heap. Are you surprised about what escapes?

package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	age int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName: lastName,
		age: age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		firstName: firstName,
		lastName: lastName,
		age: age,
	}
}

func main() {
	person1 := MakePerson("John", "Doe", 25)
	person2 := MakePersonPointer("John", "Doe", 25)
	fmt.Println(person1, person2)
}

// output of 'go build -gcflags="-m"'

// ./main.go:19:6: can inline MakePerson
// ./main.go:27:6: can inline MakePersonPointer
// ./main.go:36:23: inlining call to MakePerson
// ./main.go:37:30: inlining call to MakePersonPointer
// ./main.go:38:13: inlining call to fmt.Println
// ./main.go:19:17: leaking param: firstName to result ~r0 level=0
// ./main.go:19:35: leaking param: lastName to result ~r0 level=0
// ./main.go:27:24: leaking param: firstName
// ./main.go:27:42: leaking param: lastName
// ./main.go:28:9: &Person{...} escapes to heap
// ./main.go:37:30: &Person{...} escapes to heap
// ./main.go:38:13: ... argument does not escape
// ./main.go:38:14: person1 escapes to heap
