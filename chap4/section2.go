package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func section2a(filename string) {
	// Créer un nouveau fichier, ou le tronquer s'il existe
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Écrire des bytes dans le fichier
	bytesWritten, err := file.WriteString("Bonjour tout le monde !")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Écrit %d bytes.\n", bytesWritten)
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

func ReadNonEmptyLine(nAttempts int) (string, error) {
	for i := 0; i < nAttempts; i++ {
		line, err := ReadLine()
		if err != nil {
			fmt.Println("Erreur lors de la lecture. Essayez de nouveau.")
		} else {
			if line == "" {
				fmt.Println("La ligne ne peut pas être vide. Essayez de nouveau.")
			} else {
				return line, nil
			}
		}
	}

	return "", fmt.Errorf("échec : nombre limite de %d essai(s) atteint", nAttempts)
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

func section2b() {
	nAttempts := 3
	// demander le nom du fichier
	fmt.Print("Nom du fichier dans lequel enregistrer les données : ")
	filename, err := ReadNonEmptyLine(nAttempts)
	if err != nil {
		log.Fatalf("Impossible de lire le nom du fichier : %s", err)
	}

	// ouvrir le fichier
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Impossible d'ouvrir le fichier : %s", err)
	}
	defer file.Close()

	done := false
	for !done {
		// lire le nom
		fmt.Print("Veuillez entrer votre nom : ")
		name, err := ReadNonEmptyLine(nAttempts)
		if err != nil {
			log.Printf("Impossible de lire le nom : %s", err)
		} else {
			// s'il n'y a pas d'erreur, lire l'âge
			fmt.Print("Veuillez entrer votre âge : ")
			age, err := ReadInt(nAttempts)
			if err != nil {
				log.Printf("Impossible de lire l'âge : %s", err)
			} else {
				// s'il n'y a pas d'erreur
				line := fmt.Sprintf("%s,%d\n", name, age)
				_, err = file.WriteString(line)
				if err != nil {
					log.Fatalf("Impossible d'écrire dans le fichier : %s", err)
				}
			}
		}

		// continuer ou non ?
		fmt.Println("Voulez-vous entrer continuer (O/N) ?")
		continueInput, err := ReadNonEmptyLine(nAttempts)
		if err != nil {
			log.Printf("Impossible de lire la réponse : %s", err)
		} else {
			if strings.ToUpper(continueInput) == "N" {
				done = true
			}
		}
	}
}

func section2c() {
	nAttempts := 3
	// demander le nom du fichier
	fmt.Print("Nom du fichier dans lequel enregistrer les données : ")
	filename, err := ReadNonEmptyLine(nAttempts)
	if err != nil {
		log.Fatalf("Impossible de lire le nom du fichier : %s", err)
	}

	// ouvrir le fichier
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Impossible d'ouvrir le fichier : %s", err)
	}
	defer file.Close()

	var builder strings.Builder
	builder.WriteString("<html><body><ol>")
	done := false
	for !done {
		// lire le nom
		fmt.Print("Veuillez entrer votre nom : ")
		name, err := ReadNonEmptyLine(nAttempts)
		if err != nil {
			log.Printf("Impossible de lire le nom : %s", err)
		} else {
			builder.WriteString("<li><dl><dt>Nom</dt><dd>")
			builder.WriteString(name)
			builder.WriteString("</dd>")
			// s'il n'y a pas d'erreur, lire l'âge
			fmt.Print("Veuillez entrer votre âge : ")
			age, err := ReadInt(nAttempts)
			if err != nil {
				log.Printf("Impossible de lire l'âge : %s", err)
			} else {
				// s'il n'y a pas d'erreur
				builder.WriteString("<dt>Age</dt><dd>")
				builder.WriteString(strconv.Itoa(age))
				builder.WriteString("</dd></dl></li>")
			}
		}
		// continuer ou non ?
		fmt.Println("Voulez-vous entrer continuer (O/N) ?")
		continueInput, err := ReadNonEmptyLine(nAttempts)
		if err != nil {
			log.Printf("Impossible de lire la réponse : %s", err)
		} else {
			if strings.ToUpper(continueInput) == "N" {
				done = true
			}
		}
	}
	builder.WriteString("</ol></body></html>")
	_, err = file.WriteString(builder.String())
	if err != nil {
		log.Fatalf("Impossible d'écrire dans le fichier : %s", err)
	}
}

func readCSVFile() {
	f, err := os.Open("fichiers/abc.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Println(record) // record est une tranche de chaînes de caractères représentant les lignes du fichier
	}
}

func writeCSVFile() {
	f, err := os.Create("fichiers/def.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	data := []string{"Col1", "Col2", "Col3"}

	if err := writer.Write(data); err != nil {
		log.Fatal(err)
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
