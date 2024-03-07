// Write a function called fileLen that has an input parameter of type string
// and returns an int and an error. The function takes in a filename and
// returns the number of bytes in the file. If there is an error reading the
// file, return the error. Use defer to make sure the file is closed
// properly.

package main

import (
	"fmt"
	"os"
)

func fileLen(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return info.Size(), err
}

func main() {
	size, err := fileLen("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("example.txt is %d bytes long.\n", size)
}
