package main

import (
	"fmt"
)

func ex1() {
	var name string
	var age int

	fmt.Print("Veuillez entrer votre nom : ")
	_, err := fmt.Scanf("%s", &name)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Veuillez entrer votre âge : ")
	_, err = fmt.Scanf("%d", &age)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Bonjour %s, vous avez %d ans.\n", name, age)
}

func ex2() {
	var name string
	var age int

	fmt.Print("Veuillez entrer votre nom suivi de votre âge, séparés par un espace : ")
	n, err := fmt.Scanf("%s %d", &name, &age)

	if err != nil {
		fmt.Println(err)
	}

	if n != 2 {
		fmt.Println("Vous devez entrer votre nom suivi de votre, séparés par un espace")
		return
	}

	fmt.Printf("Bonjour %s, vous avez %d ans.\n", name, age)
}
