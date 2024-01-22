package main

import "fmt"

func swap3(xPtr *int, yPtr *int) {
	temp := *xPtr
	*xPtr = *yPtr
	*yPtr = temp
}

func main() {
	fmt.Println("Chapitre 2")

	fmt.Println("\n\n****************** Section 1 ******************")
	x := 10
	ptr := &x
	fmt.Println(*ptr)
	*ptr = 20
	fmt.Println(x)
	fmt.Println(ptr)

	fmt.Println("\n\n****************** Section 2 ******************")
	a, b := 2, 5
	fmt.Println(a, b)
	swap3(&a, &b)
	fmt.Println(a, b)

	fmt.Println("\n\n****************** Section 3 ******************")
	ex1()
	ex2()

	fmt.Println("\n\n****************** Section 4 ******************")
	section4a()
	section4b()

	// Déclaration de variable
	var name string

	// Inviter l'utilisateur à entrer son nom
	fmt.Print("Enter your name: ")

	// lire le nom depuis l'entrée standard
	_, err := fmt.Scanln(&name)

	// Vérifier s'il y a une erreur
	if err != nil {
		fmt.Println("Error reading name:", err)
	}

	// Afficher le nom
	fmt.Printf("Hello, %s!\n", name)

}
