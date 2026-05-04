package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}

type ArrayList struct {
	data []int
	len  int
}

// Add a value to the end of the array. If length equal capacity, allocate a larger
// backing slice, copyt the existing values, then store the new value.
func (a *ArrayList) Append(val int) {
	fmt.Println("Append not implemented yet")
}

// Add a value at the specific position. Check bounds, grow if needed, shift values from that
// index to the right, then place the new value.
func (a *ArrayList) Insert(val, pos int) {
	fmt.Println("Insert not implemented yet")
}

// Remove the value at a specific position. Check bounds, save the removed value if returning
// it, shift later values left, lear the unused final slot, then decrement length.
func (a *ArrayList) Remove(pos int) int {
	fmt.Println("Remove not implemlented yet")
	return 0
}

// Return the value at the specific position after checking that the index is within
// the current length.
func (a *ArrayList) Get(pos int) int {
	fmt.Println("Get not implemented yet")
	return 0
}

// Replace the value at a specific position after checking that the index is within
// the current length.
func (a *ArrayList) Set(val, pos int) {
	fmt.Println("Set not implemented yet")
}

// Return the number of values currently stored
func (a *ArrayList) Len(val, pos int) {
	fmt.Println("Len not implemented yet")
}

// Return the amount of allocated space available before another resize is required
func (a *ArrayList) Cap(val, pos int) {
	fmt.Println("Cap not implemented yet")
}
