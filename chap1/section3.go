package main

import "fmt"

func print1() {
	year := 2024
	output := fmt.Sprintln("année =", year)
	fmt.Println(output)
}

func print2() {
	year := 2024
	month := 1
	day := 5
	fmt.Printf("année = %d\n", year)
	fmt.Printf("%d-%d-%d\n", year, month, day)
	fmt.Printf("%4d-%2d-%2d\n", year, month, day)
	fmt.Printf("%4d-%02d-%02d\n", year, month, day)
}

func print3() {
	x := 1234.5678
	fmt.Printf("%f\n%.2f\n%.f\n", x, x, x)
}
