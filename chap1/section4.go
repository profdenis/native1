package main

import "fmt"

func allo(name string) {
	fmt.Printf("Allo %s !\n", name)
}

func square(x int) int {
	return x * x
}

func swap(x int, y int) (int, int) {
	return y, x
}

// cette fonction ne fonctionnera pas correctement, elle va échanger les valeurs seulement localement dans la fonction;
// il n'y aura aucun effet en dehors de la fonction
func swap2(x int, y int) {
	temp := x
	x = y
	y = temp
}
