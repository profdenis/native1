package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

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

func ReadFloat64() (float64, error) {
	line, err := ReadLine()
	if err != nil {
		return 0, err
	}
	x, err := strconv.ParseFloat(line, 64)
	if err != nil {
		return 0, err
	}
	return x, nil
}
