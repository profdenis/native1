package main

import (
	"fmt"
	"time"
)

func apply(slice []int, f func(int) int) []int {
	results := make([]int, 0, len(slice))
	for _, x := range slice {
		results = append(results, f(x))
	}
	return results
}

func plus2(x int) int {
	return x + 2
}

func newAdder(x int) func(int) int {
	return func(i int) int {
		return i + x
	}
}

func printNumbers(x int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", x*10+i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printNumbersChan(x int, ch chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", x*10+i)
		time.Sleep(100 * time.Millisecond)
	}
	ch <- x
}

func main() {
	//numbers := []int{4, 8, 1, 9, 0, -3, 5}
	//fmt.Println(numbers)
	//results := apply(numbers, plus2)
	//fmt.Println(results)
	//results = apply(numbers, func(i int) int {
	//	return i * 2
	//})
	//fmt.Println(results)
	//plus5 := newAdder(5)
	//results = apply(numbers, plus5)
	//fmt.Println(results)
	//fmt.Println(plus5(7))

	//for i := 0; i < 5; i++ {
	//	fmt.Printf("-%d- ", i)
	//	go printNumbers(i)
	//}
	//time.Sleep(1 * time.Second)

	ch := make(chan int)
	for i := 0; i < 5; i++ {
		fmt.Printf("-%d- ", i)
		go printNumbersChan(i, ch)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
