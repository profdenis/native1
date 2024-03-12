package main

import "fmt"

func sumSlice(ints []int) int {
	sum := 0
	for _, num := range ints {
		sum += num
	}
	return sum
}

func section3a() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("La somme des numéros est : %d\n", sumSlice(nums))
}

func sumSliceChan(ints []int, results chan<- int) {
	sum := 0
	for _, num := range ints {
		sum += num
	}
	results <- sum
}

func parallelSum(nums []int) int {
	// Création du canal
	results := make(chan int)

	// Divise la slice et lance les goroutines.
	half := len(nums) / 2
	go sumSliceChan(nums[:half], results)
	go sumSliceChan(nums[half:], results)

	// Recevoir les résultats des goroutines et calcule la somme totale.
	sum1, sum2 := <-results, <-results

	return sum1 + sum2
}

func section3b() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	totalSum := parallelSum(nums)
	fmt.Printf("La somme des numéros est : %d\n", totalSum)
}
