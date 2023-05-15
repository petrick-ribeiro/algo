package main

import "fmt"

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue adds a value at the end of the slice
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes a value at the beginning of the slice
// and returns the removed value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0] // Get the first value
	q.items = q.items[1:]

	return toRemove
}

func main() {
	q := Queue{}
	fmt.Println(q.items)

	foo := []int{2, 42, 0, 101}
	for _, v := range foo {
		q.Enqueue(v)
		fmt.Println(q.items)
	}

	for i := 0; i < len(foo); i++ {
		q.Dequeue()
		fmt.Println(q.items)
	}
}
