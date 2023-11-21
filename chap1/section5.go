package main

import "fmt"

func conditional1(x int) {
	if x > 0 {
		fmt.Printf("x = %d est plus grand que 0\n", x)
	} else {
		fmt.Printf("x = %d est plus petit ou égal à 0\n", x)
	}
}

func conditional2(x int) {
	if x > 0 {
		fmt.Printf("x = %d est plus grand que 0\n", x)
	} else if x == 0 {
		fmt.Println("x est égal à 0")
	} else {
		fmt.Printf("x = %d est plus petit que 0\n", x)
	}
}

func conditional3(x int) {
	switch {
	case x > 0:
		fmt.Printf("x = %d est plus grand que 0\n", x)
	case x == 0:
		fmt.Println("x est égal à 0")
	case x < 0:
		fmt.Printf("x = %d est plus petit que 0\n", x)
	}
}

func conditional4(x int) {
	switch {
	case x > 0:
		fmt.Printf("x = %d est plus grand que 0\n", x)
	case x == 0:
		fmt.Println("x est égal à 0")
	default:
		fmt.Printf("x = %d est plus petit que 0\n", x)
	}
}

func conditional5(name string) {
	switch name {
	case "Denis":
		fmt.Println("Bonjour Prof. Denis !")
	case "":
		fmt.Println("Bonjour Inconnu !")
	default:
		fmt.Printf("Allo %s !\n", name)
	}
}

func conditional6(name string) {
	switch name {
	case "Denis", "Benoit":
		fmt.Printf("Bonjour Prof. %s !\n", name)
	case "":
		fmt.Println("Bonjour Inconnu !")
	default:
		fmt.Printf("Allo %s \n!", name)
	}
}
