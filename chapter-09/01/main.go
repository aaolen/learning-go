// Create a sentinel error to represent an invalid ID. In main, use errors.Is
// to check for the sentinel error, and print a message when it is found.

// Define a custom error type to represent an empty field error. This error
// should include the name of the empty Employee field. In main, use errors.As
// to check for this error. Print out a message that includes the field name.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var ErrInvalidID error = errors.New("invalid id")

type EmptyFieldError struct {
	field string
}

func (e EmptyFieldError) Error() string {
	return e.field
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		if err != nil {
			var errEmptyField EmptyFieldError
			if errors.Is(err, ErrInvalidID) {
				fmt.Printf("record %d: %+v invalid id: %v\n", count, emp, emp.ID)
			} else if errors.As(err, &errEmptyField) {
				fmt.Printf("record %d: %+v missing field: %v\n", count, emp, err)
			} else {
				fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
			}
			continue
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

func ValidateEmployee(e Employee) error {
	if len(e.ID) == 0 {
		return EmptyFieldError{"ID"}
	}
	if !validID.MatchString(e.ID) {
		return ErrInvalidID
	}
	if len(e.FirstName) == 0 {
		return EmptyFieldError{"FirstName"}
	}
	if len(e.LastName) == 0 {
		return EmptyFieldError{"LastName"}
	}
	if len(e.Title) == 0 {
		return EmptyFieldError{"Title"}
	}
	return nil
}
