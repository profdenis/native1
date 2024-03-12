package main

import "fmt"

// Cette fonction crée une fermeture
func createAdder(x int) func(int) int {
	// La fonction retournée est une fermeture
	return func(y int) int {
		// Elle "se souvient" de la valeur de x
		return x + y
	}
}

func section1a() {
	adder := createAdder(10)
	fmt.Println(adder(5)) // Affiche 15
}

// Notre fonction prend une tranche et une fonction en paramètre
func apply(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = fn(num)
	}
	return result
}

func section1b() {
	nums := []int{1, 2, 3, 4, 5}

	// Nous passons une fonction anonyme qui multiplie chaque nombre par 2
	doubled := apply(nums, func(n int) int {
		return n * 2
	})

	fmt.Println(doubled) // Affiche : [2 4 6 8 10]
}

func multiplyByTwo(n int) int {
	return n * 2
}

func printNumber(n int) {
	fmt.Println(n)
}

func section1c() {
	// La fonction multiplyByTwo est appelée immédiatement avec 10 comme argument
	// printNumber reçoit uniquement le résultat de cette fonction, soit 20.
	printNumber(multiplyByTwo(10))
}

func printAfterOperation(n int, fn func(int) int) {
	fmt.Println(fn(n))
}

func section1d() {
	// Ici, multiplyByTwo est passé en tant que référence de fonction à printAfterOperation
	// Il ne sera exécuté qu'à l'intérieur de printAfterOperation, avec 10 comme argument
	printAfterOperation(10, multiplyByTwo)
}
