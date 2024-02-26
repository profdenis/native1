package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func f(i int) int {
	if i <= 0 {
		return 0
	}
	return i + f(i-1)
}

func fib(i int) int {
	if i <= 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	a := 0
	b := 1
	for j := 2; j <= i; j++ {
		a, b = b, a+b
	}
	return b
}

func section1() {
	x, err := ReadInt(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}

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
