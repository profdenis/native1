package main

import "fmt"

func main() {
	// section 1
	fmt.Println("\n\n****************** Section 1 ******************")
	fmt.Println("Hello World!")

	// section 2
	fmt.Println("\n\n****************** Section 2 ******************")
	variables1()
	variables2()
	variables3()
	variables4()

	// section 3
	fmt.Println("\n\n****************** Section 3 ******************")
	print1()
	print2()
	print3()

	// section 4
	fmt.Println("\n\n****************** Section 4 ******************")
	allo("Alice")
	fmt.Println(square(4))

	a, b := 2, 5
	fmt.Println(swap(a, b))

	fmt.Println(a, b)
	swap2(a, b)
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)

	// section 5
	fmt.Println("\n\n****************** Section 5 ******************")
	funcs := []func(int){conditional1, conditional2, conditional3, conditional4}
	params := []int{5, 0, -5}
	for _, f := range funcs {
		for _, x := range params {
			f(x)
		}
	}
	conditional5("Denis")
	conditional5("")
	conditional5("Alice")

	conditional6("Denis")
	conditional6("Benoit")

	// section 6
	fmt.Println("\n\n****************** Section 6 ******************")
	loop1()
	loop2()
	loop3()
	loop4()
	loop5()
	loop6()

}
