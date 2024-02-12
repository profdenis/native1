package main

import "fmt"

//type Person struct {
//	Name string
//	Age  int
//}

func main() {

	fmt.Println("\n\n****************** Section 1 ******************")
	// Initialiser une nouvelle instance de Person
	person1 := Person{Name: "John Doe", Age: 30}

	// Accéder aux champs de la struct
	fmt.Println(person1.Name) // John Doe
	fmt.Println(person1.Age)  // 30

	// Modifier les champs de la struct
	person1.Name = "Jane Doe"
	person1.Age = 28

	fmt.Println(person1.Name) // Jane Doe
	fmt.Println(person1.Age)  // 28

	fmt.Println(person1)

	section1a()

	fmt.Println("\n\n****************** Section 2 ******************")
	section2a()

	fmt.Println("\n\n****************** Section 3 ******************")
	section3a()
	section3b()
	section3c()
}
