package main

import (
	"log"
	"os"
)

func section4a() {
	// sans os.O_CREATE, le fichier ne sera pas créé s'il n'existe pas déjà
	// avec os.O_CREATE, le fichier sera créé seulement s'il n'existe pas déjà
	// dans les 2 cas, le fichier ne sera pas tronqué, on va écrire par-dessus le contenu existant
	// il faut utiliser os.O_TRUNC pour tronquer le contenu
	file, err := os.OpenFile("file.txt", os.O_WRONLY, 0644)
	//file, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	//file, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, err = file.WriteString("Section4a\n")
	if err != nil {
		return
	}
}

func section4b() {
	f, err := os.OpenFile("file2.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	_, err = f.WriteString("Section4b\n")
}

func section4c() {
	file, err := os.OpenFile("file3.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, err = file.WriteString("Section4c\n")
	if err != nil {
		return
	}
	err = file.Sync()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = file.Seek(40, 0)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = file.WriteString("Section4c\n")
	if err != nil {
		return
	}
}
