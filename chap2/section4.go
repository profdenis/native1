package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

//func ReadLine() (string, error) {
//	reader := bufio.NewReader(os.Stdin)
//	text, err := reader.ReadString('\n')
//	if err != nil {
//		return "", err
//	} else {
//		return text[:len(text)-1], nil
//	}
//}

func ReadLine() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text(), nil
	} else {
		err := scanner.Err()
		if err != nil {
			return "", err
		} else {
			return "", io.EOF
		}
	}
}

func section4a() {
	fmt.Print("Veuillez entrer votre nom : ")
	name, err := ReadLine()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Veuillez entrer votre âge : ")
	ageStr, err := ReadLine()

	if err != nil {
		fmt.Println(err)
	}

	age, err := strconv.Atoi(ageStr)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Bonjour %s, vous avez %d ans.\n", name, age)
}

func ReadInt(nAttempts int) (int, error) {
	for i := 0; i < nAttempts; i++ {
		inputStr, err := ReadLine()
		if err != nil {
			fmt.Println("Erreur lors de la lecture. Essayez de nouveau.")
		} else {
			inputNum, err := strconv.Atoi(inputStr)
			if err != nil {
				fmt.Println("Ceci n'est pas un nombre entier. Essayez de nouveau.")
			} else {
				return inputNum, nil
			}
		}
	}

	return 0, fmt.Errorf("échec : nombre limite de %d essai(s) atteint", nAttempts)
}

func section4b() {
	fmt.Print("Veuillez entrer votre nom : ")
	name, err := ReadLine()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Veuillez entrer votre âge : ")
	age, err := ReadInt(3)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Bonjour %s, vous avez %d ans.\n", name, age)
}
