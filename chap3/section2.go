package main

import "fmt"

func modSlice1(s []int) {
	s[0] = 100
}

func modSlice2(s []int) {
	s = append(s, 100)
}

func modSlice3(s *[]int) {
	*s = append(*s, 100)
}

func modSlice4(s []int) {
	s = append(s, 100)
}

func modSlice5(s *[]int) {
	*s = append(*s, 100)
}

func section2a() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	modSlice1(s)
	fmt.Println(s)
	fmt.Println()
}

func section2b() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	modSlice2(s)
	fmt.Println(s)
	fmt.Println()
}

func section2c() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	modSlice3(&s)
	fmt.Println(s)
	fmt.Println()
}

func section2d() {
	s := make([]int, 3, 5)
	s[0] = 4
	s[1] = 1
	s[2] = -5
	fmt.Printf("%d %d %v\n", len(s), cap(s), s)
	modSlice4(s)
	fmt.Printf("%d %d %v\n\n", len(s), cap(s), s)
}

func section2e() {
	s := make([]int, 3, 5)
	s[0] = 4
	s[1] = 1
	s[2] = -5
	fmt.Printf("%d %d %v\n", len(s), cap(s), s)
	modSlice5(&s)
	fmt.Printf("%d %d %v\n", len(s), cap(s), s)
	modSlice5(&s)
	modSlice5(&s)
	fmt.Printf("%d %d %v\n\n", len(s), cap(s), s)

}

func section2f() {
	numbers := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(numbers[1:6])
	fmt.Println(numbers[1:5])
	sub := numbers[1:]
	fmt.Println(sub)
	numbers[2] = -1
	fmt.Println(numbers)
	fmt.Println(sub)
	fmt.Println(sub[:len(sub)-1])
	fmt.Println(sub[1:2])
	fmt.Println(sub[1])
}
