// Write a program that builds a []Person with 10,000,000 entries (they could
// all be the same names and ages). See how long it takes to run. Change the
// value of GOGC and see how that affects the time it takes for the program to
// complete. Set the environment variable GODEBUG=gctrace=1 to see when
// garbage collections happen and see how changing GOGC changes the number of
// garbage collections. What happens if you create the slice with a capacity
// of 10,000,000?

package main

import (
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName string
	age int
}

func main() {
	start := time.Now()
	people := make([]Person, 10_000_000)
	for i := range people {
		people[i] = Person{
			firstName: "John",
			lastName: "Doe",
			age: 25,
		}
	}
	executionTime := time.Since(start)
	fmt.Println(len(people), executionTime)
}
