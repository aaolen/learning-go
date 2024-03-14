// Define a generic interface called Printable that matches a type that
// implements fmt.Stringer and has an underlying type of int or float64.
// Define types that meet this interface. Write a function that takes in a
// Printable and prints its value to the screen using fmt.Println.

package main

import "fmt"

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type PrintableInt int

func (pi PrintableInt) String() string {
	return fmt.Sprintf("%d", pi)
}

type PrintableFloat float64

func (pf PrintableFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func Print[T Printable](p T) {
	fmt.Println(p)
}

func main() {
	var i PrintableInt = 32
	var f PrintableFloat = 24.5
	Print(i)
	Print(f)
}
