package main

import (
	"bufio"
	"fmt"
	"os"
)

func CopyFileWithLineNumbers(srcFilename string, dstFilename string) error {
	// Ouvrez le fichier source en lecture seule
	src, err := os.Open(srcFilename)
	if err != nil {
		return err
	}
	defer src.Close()

	// Ouvrez le fichier de destination en mode écriture
	dst, err := os.Create(dstFilename)
	if err != nil {
		return err
	}
	defer dst.Close()

	scanner := bufio.NewScanner(src)
	writer := bufio.NewWriter(dst)
	lineNumber := 1

	for scanner.Scan() {
		lineStr := fmt.Sprintf("%d\t%s\n", lineNumber, scanner.Text())
		if _, err := writer.WriteString(lineStr); err != nil {
			return err
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

func section3() {
	err := CopyFileWithLineNumbers("fichiers/allo.txt", "fichiers/allo2.txt")
	if err != nil {
		fmt.Println("Erreur lors de la copie du fichier:", err)
	} else {
		fmt.Println("Le fichier a été copié avec succès avec des numéros de ligne.")
	}
}
