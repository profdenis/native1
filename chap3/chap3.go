package main

import (
	"fmt"
	"strings"
)

// La fonction modifierTentative modifie le tableau qu'elle reçoit.
func modifierTentative(arr [3]int) {
	arr[0] = 100
}

func main() {
	monTableau := [3]int{1, 2, 3}

	modifierTentative(monTableau)

	fmt.Println(monTableau) // Sortie : [1 2 3]

	autreTableau := monTableau
	autreTableau[0] = 100

	fmt.Println(monTableau) // Sortie : [1 2 3]

	//a := "abc"
	//a[1] = 'x' // ne compilera pas

	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	builder.WriteString("!")

	result := builder.String()

	fmt.Println(result)
}
