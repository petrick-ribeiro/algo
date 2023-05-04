package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	prevToDelete := l.head

	for prevToDelete.next.data != value {
		if prevToDelete.next.next == nil {
			return
		}

		prevToDelete = prevToDelete.next
	}

	prevToDelete.next = prevToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	nodes := []int{42, 35, 14, 99, 1, 21}

	for _, value := range nodes {
		myList.prepend(&node{data: value})
	}

	myList.printListData()
	myList.deleteWithValue(99)
	myList.printListData()
}
