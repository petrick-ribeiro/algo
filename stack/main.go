package main

import "fmt"

// Stack represents a stack that contains a slice.
type Stack struct {
	items []int
}

// Push add a value at the end of the slice.
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop remove a value at the end of the slice
// and returns the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1 // Get the last index of the slice.
	toRemove := s.items[l]
	s.items = s.items[:l]

	return toRemove
}

func main() {
	s := Stack{}
	fmt.Println(s.items)

	foo := []int{14, 1, 101, 42}
	for _, v := range foo {
		s.Push(v)
		fmt.Println(s.items)
	}

	for i := 0; i < len(foo); i++ {
		s.Pop()
		fmt.Println(s.items)
	}
}
