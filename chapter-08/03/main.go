// Write a generic singly linked list data type. Each element can hold a
// comparable value and has a pointer to the next element in the list. The
// methods to implement are as follows:
//
// // adds a new element to the end of the linked list
// Add(T)
// // adds an element at the specified position in the linked list
// Insert(T, int)
// // returns the position of the supplied value, -1 if it's not present
// Index (T) int

package main

import (
	"errors"
	"fmt"
	"os"
)

type ComparableStringer interface {
	fmt.Stringer
	comparable
}

type Node[T ComparableStringer] struct {
	data T
	next *Node[T]
	fmt.Stringer
}

type LinkedList[T ComparableStringer] struct {
	head *Node[T]
	fmt.Stringer
}

func  NewLinkedList[T ComparableStringer](t T) *LinkedList[T] {
	h := Node[T]{ data: t }
	return &LinkedList[T]{
		head: &h,
	}
}

func (l LinkedList[T]) String() string {
	s := ""
	if l.head == nil {
		return "[]"
	}
	current := l.head
	for current != nil {
		if l.head != current {
			s += ", "
		}
		s += fmt.Sprint(current.data)
		current = current.next
	}
	return fmt.Sprintf("[%s]", s)
}

func (l *LinkedList[T]) Add(t T) {
	newNode := Node[T]{ data: t }
	if l.head == nil {
		l.head = &newNode
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &newNode
}

func (l *LinkedList[T]) Insert(t T, p int) error {
	newNode := Node[T]{ data: t }
	if p == 0 {
		newNode.next = l.head
		l.head = &newNode
		return nil
	}
	current := l.head
	i := 1
	for current != nil {
		if i == p {
			newNode.next = current.next
			current.next = &newNode
			return nil
		}
		i++
		current = current.next
	}
	return errors.New("index out of range")
}

func (l *LinkedList[T]) Index(t T) int {
	current := l.head
	i := 0
	for current != nil {
		if current.data == t {
			return i
		}
		current = current.next
		i++
	}
	return -1
}

type Person struct {
	fmt.Stringer
	firstName string
	lastName string
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func main() {
	myLinkedList := NewLinkedList(Person{
		firstName: "John",
		lastName: "Doe",
	})
	myLinkedList.Add(Person{
		firstName: "Johnny",
		lastName: "Appleseed",
	})
	myLinkedList.Add(Person{
		firstName: "Bob",
		lastName: "Smith",
	})
	err := myLinkedList.Insert(Person{
		firstName: "Sam",
		lastName: "Smith",
	}, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Result: %s\n", myLinkedList)
	fmt.Printf("Index of 'Sam Smith': %d\n", myLinkedList.Index(Person{
		firstName: "Sam",
		lastName: "Smith",
	}))
}
