package main

import (
	"fmt"
)

func main() {
	done := false
	for !done {
		showMenu()
		option, _ := ReadLine()
		done = handleOption(option)
	}
}

func showMenu() {
	fmt.Println("Choisissez une option :")
	fmt.Println("1. Option 1")
	fmt.Println("2. Option 2")
	fmt.Println("3. Option 3")
	fmt.Println("4. Quitter")
}

func handleOption(option string) bool {
	switch option {
	case "1":
		fmt.Println("Vous avez choisi l'option 1 !")
	case "2":
		fmt.Println("Vous avez choisi l'option 2 !")
	case "3":
		fmt.Println("Vous avez choisi l'option 3 !")
	case "4":
		fmt.Println("Vous voulez quitter !")
		return true
	default:
		fmt.Println("Option invalide. SVP choisir une option valide.")
	}
	return false
}
