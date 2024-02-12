package main

import "fmt"

func section2a() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(3, 4)
	fmt.Println(p1)
	fmt.Println(p2)

	p3 := p1.Add(p2)
	fmt.Println(p3)
}
