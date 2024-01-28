package main

import "fmt"

func ex01() {
	fmt.Print("Entrez un nombre : ")
	x, _ := ReadFloat64()
	if x < 0 {
		fmt.Printf("La valeur absolue de x est %f\n", -x)
	} else {
		fmt.Printf("La valeur absolue de x est %f\n", x)
	}
}

func ex02() {
	fmt.Print("Entrez un nombre : ")
	x, _ := ReadFloat64()
	i := int(x)
	if i%2 == 0 {
		fmt.Printf("Le nombre %d est pair.\n", i)
	} else {
		fmt.Printf("Le nombre %d est impair.\n", i)
	}
}

func ex03() {
	fmt.Print("Entrez un premier nombre : ")
	x, _ := ReadFloat64()
	fmt.Print("Entrez un deuxième nombre : ")
	y, _ := ReadFloat64()
	fmt.Print("Entrez un troisième nombre : ")
	z, _ := ReadFloat64()
	if x < y && x < z {
		fmt.Printf("Le nombre %f est le plus petit des trois nombres.\n", x)
	} else if y < x && y < z {
		fmt.Printf("Le nombre %f est le plus petit des trois nombres.\n", y)
	} else if z < x && z < y {
		fmt.Printf("Le nombre %f est le plus petit des trois nombres.\n", z)
	}
}

func ex04() {
	fmt.Print("Entrez le salaire horaire : ")
	salaire, _ := ReadFloat64()
	fmt.Print("Entrez le nombre d'heures : ")
	heures, _ := ReadFloat64()

	var total float64

	if heures > 40 {
		total = salaire * 40
		total += salaire * (heures - 40) * 1.5
	} else {
		total = salaire * heures
	}
	fmt.Printf("Le salaire total est %.2f$", total)
}

func ex05() {
	fmt.Print("Entrez la longueur du premier côté : ")
	x, _ := ReadFloat64()
	fmt.Print("Entrez la longueur du deuxième côté : ")
	y, _ := ReadFloat64()
	fmt.Print("Entrez la longueur du troisième côté : ")
	z, _ := ReadFloat64()

	if x == y && y == z {
		fmt.Println("Équilatéral")
	} else if x != y && y != z && x != z {
		fmt.Println("Scalène")
	} else {
		fmt.Println("Isocèle")
	}
}

func ex06() {
	fmt.Print("Entrez une note : ")
	total, _ := ReadFloat64()
	fmt.Print("Entrez une note : ")
	x, _ := ReadFloat64()
	total += x
	fmt.Print("Entrez une note : ")
	x, _ = ReadFloat64()
	total += x
	avg := total / 3

	if avg < 60 {
		fmt.Println("Échec")
	} else {
		fmt.Println(avg)
	}
}

func ex07() {
	fmt.Print("Entrez une note : ")
	score, _ := ReadFloat64()
	//if score < 0 || score > 100 {
	//	fmt.Println("Cette note est invalide")
	//	return
	//}
	switch {
	case score < 0 || score > 100:
		fmt.Println("Cette note est invalide")
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}

func ex08() {
	fmt.Print("Entrez un nombre : ")
	x, _ := ReadFloat64()
	//if x < 1 || x > 10 {
	//	fmt.Println("invalide")
	//} else {
	//	fmt.Println("valide")
	//}
	if x >= 1 && x <= 10 {
		fmt.Println("valide")
	} else {
		fmt.Println("invalide")
	}
}

func ex09() {
	valid := false
	for !valid {
		fmt.Print("Entrez un nombre : ")
		x, _ := ReadFloat64()
		if x >= 1 && x <= 10 {
			fmt.Println("valide")
			valid = true
		} else {
			fmt.Println("invalide")
		}
	}
}

func ex10() {
	valid := false
	n := 0
	for !valid && n < 3 {
		fmt.Print("Entrez un nombre : ")
		x, _ := ReadFloat64()
		n++
		if x >= 1 && x <= 10 {
			valid = true
		} else {
			fmt.Println("invalide")
		}
	}
	if valid {
		fmt.Println("valide")
	} else {
		fmt.Println("Nombre maximal d'essais atteint.")
	}
}

func ex11() {
	fmt.Print("Entrez un nombre : ")
	x, _ := ReadFloat64()
	i := int(x)
	var j int

	for i > 0 {
		//fmt.Print(i % 10)
		j = j*10 + i%10
		i = i / 10
	}

	//fmt.Println()
	fmt.Println(j)
}

func ex12() {
	fmt.Print("Entrez un nombre : ")
	x, _ := ReadFloat64()
	i := int(x)

	for i > 0 {
		fmt.Println(i)
		i--
	}
	fmt.Println("Terminé !")
}

func ex13a() {
	numbers := [...]int{2, 5, 3, 4, 7, 1, 7}
	var total int
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	avg := float64(total) / float64(len(numbers))
	fmt.Printf("La somme est %d et le moyenne est %f\n", total, avg)
}

func ex13b() {
	numbers := [...]int{2, 5, 3, 4, 7, 1, 7}
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			fmt.Println(numbers[i])
		}
	}

}

func ex13c() {
	numbers := [...]int{2, 5, 3, 4, 7, 1, 7}
	pos := true
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < 0 {
			pos = false
			break
		}
	}
	if pos {
		fmt.Println("vrai")
	} else {
		fmt.Println("faux")
	}
}

func ex14a() {
	numbers := [...]int{2, 5, 3, 4, 7, 1, 7}
	var total int
	for _, number := range numbers {
		total += number
	}
	avg := float64(total) / float64(len(numbers))
	fmt.Printf("La somme est %d et le moyenne est %f\n", total, avg)
}

func ex14b() {
	numbers := [...]int{2, 5, 3, 4, 7, 1, 7}
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number)
		}
	}
}

func ex14c() {
	numbers := [...]int{2, 5, 3, 4, 7, 1, 7}
	pos := true

	for _, number := range numbers {
		if number < 0 {
			pos = false
			break
		}
	}
	if pos {
		fmt.Println("vrai")
	} else {
		fmt.Println("faux")
	}
}
