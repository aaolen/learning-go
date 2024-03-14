// Write a generic function that doubles the value of any integer or float
// that’s passed in to it. Define any needed generic interfaces.

package main

import "fmt"

type Numeric interface {
	~int | ~int16 | ~int32 | ~int64 | ~int8 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 
}

func DoubleNum[T Numeric](n T) T {
	return (n * 2)
}

func main() {
	var n1 int = 3
	var n2 float64 = 4.4
	var n3 float32 = 3.15
	fmt.Println(DoubleNum(n1))
	fmt.Println(DoubleNum(n2))
	fmt.Println(DoubleNum(n3))
}
