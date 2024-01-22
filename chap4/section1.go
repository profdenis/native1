package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func section1a() {
	content, err := os.ReadFile("fichiers/allo.txt") // Lecture du contenu du fichier
	if err != nil {
		log.Fatalf("lecture du contenu du fichier impossible : %s", err)
	}

	fmt.Println("Contenu du fichier :")
	fmt.Println(string(content))
}

func section1b() {
	file, err := os.Open("fichiers/allo.txt") // Ouverture du fichier en mode lecture
	if err != nil {
		log.Fatalf("ouverture du fichier impossible : %s", err)
	}
	//defer file.Close()

	content, err := io.ReadAll(file) // Lecture du contenu du fichier
	if err != nil {
		log.Fatalf("lecture du contenu du fichier impossible : %s", err)
	}

	fmt.Println("Contenu du fichier :")
	fmt.Println(string(content))

	err = file.Close()
	if err != nil {
		log.Printf("erreur lors de la fermeture du fichier : %s", err)
	}
}

func countNonEmptyLines(filename string) {
	file, err := os.Open(filename) // Ouverture du fichier en mode lecture
	if err != nil {
		log.Fatalf("ouverture du fichier impossible : %s", err)
	}

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			count++
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Printf("erreur lors de la lecture du fichier : %s", err)
	}

	fmt.Println("Nombre de ligne non-vide :", count)

	err = file.Close()
	if err != nil {
		log.Printf("erreur lors de la fermeture du fichier : %s", err)
	}
}
