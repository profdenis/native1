package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func section2() {
	// Allocation sur le tas
	pHeap := &Person{"Alice", 30}
	fmt.Printf("Address of person on heap: %p\n", pHeap)

	// Allocation sur la pile
	pStack := Person{"Bob", 35}
	fmt.Printf("Address of person on stack: %p\n", &pStack)
}
