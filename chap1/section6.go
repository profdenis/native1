package main

import "fmt"

func loop1() {
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func loop2() {
	i := 0
	for i < 5 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()
}

func loop3() {
	i := 0
	for {
		fmt.Print(i, " ")
		i++
		if i >= 5 {
			break
		}
	}
	fmt.Println()
}

func loop4() {
	numbers := [7]int{2, 5, 3, 4, 7, 1, 7}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
	}
	fmt.Println(numbers)
}

func loop5() {
	numbers := [...]int{2, 5, 3, 4, 7, 1, 7}
	for i, number := range numbers {
		fmt.Printf("numbers[%d] = %d\n", i, number)
	}
	fmt.Println(numbers)
}

func loop6() {
	numbers := [...]int{2, 5, 3, 4, 7, 1, 7}
	for _, number := range numbers {
		fmt.Printf("%d ", number)
	}
	fmt.Println()
}
