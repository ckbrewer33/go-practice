package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
	var a ArrayList

	fmt.Println("appending '1'")
	a.Append(1)

	fmt.Print("current length: ")
	fmt.Println(a.len)

	fmt.Println("appending '2'")
	a.Append(2)

	fmt.Print("current length: ")
	fmt.Println(a.len)
}

type ArrayList struct {
	data []int
	len  int
}

// Add a value to the end of the array. If length equal capacity, allocate a larger
// backing slice, copyt the existing values, then store the new value.
func (a *ArrayList) Append(val int) {
	a.data = append(a.data, val)
	a.len = a.len + 1
}

// Add a value at the specific position. Check bounds, grow if needed, shift values from that
// index to the right, then place the new value.
func (a *ArrayList) Insert(val, pos int) {
	fmt.Println("Insert not implemented yet")
}

// Remove the value at a specific position. Check bounds, save the removed value if returning
// it, shift later values left, clear the unused final slot, then decrement length.
func (a *ArrayList) Remove(pos int) int {
	fmt.Println("Remove not implemlented yet")
	return 0
}

// Return the value at the specific position after checking that the index is within
// the current length.
func (a *ArrayList) Get(pos int) (int, error) {
	if len(a.data) < 1 {
		return 0, errors.New("Array list is empty")
	}

	return a.data[pos], nil
}

// Replace the value at a specific position after checking that the index is within
// the current length.
func (a *ArrayList) Set(val, pos int) {
	fmt.Println("Set not implemented yet")
}

// Return the number of values currently stored
func (a *ArrayList) Len() int {
	return a.len
}

// Return the amount of allocated space available before another resize is required
func (a *ArrayList) Cap() {
	fmt.Println("Cap not implemented yet")
}
