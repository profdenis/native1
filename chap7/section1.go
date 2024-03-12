package main

import "fmt"

type Person struct {
	Id   int
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("ID: %d, Nom : %s, Âge : %d\n", p.Id, p.Name, p.Age)
}

func section1a() {
	// Création des personnes
	person1 := &Person{Id: 1, Name: "John Doe", Age: 30}
	person2 := &Person{Id: 2, Name: "Jane Doe", Age: 28}
	person3 := &Person{Id: 3, Name: "Richard Roe", Age: 33}
	person4 := &Person{Id: 4, Name: "Maggie Roe", Age: 32}

	// Création de la map pour stocker les personnes par leur ID
	persons := make(map[int]*Person)

	// Ajout des personnes dans la map
	persons[person1.Id] = person1
	persons[person2.Id] = person2
	persons[person3.Id] = person3
	persons[person4.Id] = person4

	// Affichage des personnes de la map
	for _, person := range persons {
		fmt.Print(person)
	}
}
