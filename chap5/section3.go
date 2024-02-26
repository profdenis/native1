package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//func (p Person) String() string {
//	return fmt.Sprintf("Person(%s, %d)", p.Name, p.Age)
//}

func section3a() {
	// Créer une instance de Person
	p := Person{Name: "John", Age: 30}

	// Sérialiser la struct en JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData)) // Affiche : {"name":"John","age":30}

	// Désérialiser le JSON en struct
	jsonStr := `{"name":"James","age":40}`
	var p2 Person
	err = json.Unmarshal([]byte(jsonStr), &p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2) // Affiche : {James 40}
}

func section3b() {
	// Créer une instance de Person
	p := Person{Name: "Denis", Age: 30}

	// Sérialiser la struct en JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Ouvrir le fichier en mode écriture
	file, err := os.OpenFile("person.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Écrire les données JSON dans le fichier
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Sync()

	// Lire les données du fichier
	data, err := os.ReadFile("person.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Désérialiser les données JSON pour recréer une instance de Person
	var p2 Person
	err = json.Unmarshal(data, &p2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Afficher l'instance de Person
	fmt.Println(p2) // Devrait afficher: {John 30}
}

func section3c() {
	// Créer une tranche de Person
	people := []Person{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Alice", Age: 35},
	}

	// Sérialiser la slice en JSON
	jsonData, err := json.MarshalIndent(people, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Ouvrir le fichier en mode écriture
	file, err := os.OpenFile("people.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Écrire les données JSON dans le fichier
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Sync()

	// Lire les données du fichier
	data, err := os.ReadFile("people.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Désérialiser les données JSON pour recréer une tranche de Person
	var people2 []Person
	err = json.Unmarshal(data, &people2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Afficher la slice de Person
	fmt.Println(people2) // Devrait afficher : [{John 30} {Jane 25} {Alice 35}]
}

type People []Person

func (people People) String() string {
	var builder strings.Builder
	for i, p := range people {
		builder.WriteString(fmt.Sprintf("%d: %v\n", i, p))
	}

	return builder.String()
}

func section3d() {
	// Créer une tranche de Person
	people := People{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Alice", Age: 35},
	}

	err := WritePeople(people, "fichiers/people.json")
	if err != nil {
		log.Fatal(err)
	}

	err = ReadPeople(&people, "fichiers/people.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(people)
}

func ReadPeople(people *People, filename string) error {
	// Lire les données du fichier
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Désérialiser les données JSON pour recréer une tranche de Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		return err
	}

	return nil
}

func WritePeople(people People, filename string) error {
	// Sérialiser la slice en JSON
	jsonData, err := json.MarshalIndent(people, "", "\t")
	if err != nil {
		return err
	}

	// Ouvrir le fichier en mode écriture
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Écrire les données JSON dans le fichier
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}
