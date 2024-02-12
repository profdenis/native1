package main

import "fmt"

func section1a() {
	p := Point{2, 5}
	fmt.Println(p)

	pPtr := &p
	fmt.Printf("%p %v\n", pPtr, *pPtr)

	pPtr = &Point{4, 3}
	fmt.Printf("%p %v\n", pPtr, *pPtr)

	pPtr = new(Point)
	fmt.Printf("%p %v\n", pPtr, *pPtr)
	pPtr.X = 1
	pPtr.Y = 6
	fmt.Printf("%p %v\n", pPtr, *pPtr)

	pPtr = NewPoint(-1, 5.5)
	fmt.Printf("%p %v\n", pPtr, *pPtr)

}
