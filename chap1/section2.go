package main

import "fmt"

func variables1() {
	var i int
	var finished bool
	var name string
	var x, y float32

	fmt.Println(i)
	fmt.Println(finished)
	fmt.Println(name)
	fmt.Println(x, y)
}

func variables2() {
	var i int = 5
	var finished bool = true
	var name string = "Denis"
	var x, y float32 = 2.5, -1.4

	fmt.Println(i)
	fmt.Println(finished)
	fmt.Println(name)
	fmt.Println(x, y)
}

func variables3() {
	var i = 5
	var finished = true
	var name = "Denis"
	var x, y = 2.5, -1.4

	fmt.Println(i)
	fmt.Println(finished)
	fmt.Println(name)
	fmt.Println(x, y)
	fmt.Printf("%T\n", x)
}

func variables4() {
	i := 5
	finished := true
	name := "Denis"
	x := 2.5
	y := -1.4

	fmt.Println(i)
	fmt.Println(finished)
	fmt.Println(name)
	fmt.Println(x, y)
	fmt.Printf("%T\n", x)
}
