package main

import (
	"fmt"
	"sort"
)

func section4a() {
	// trier une tranche de nombres entiers
	numbers1 := []int{2, 5, 1, 7, 3, 8, 4, 7, 6, 4}
	fmt.Println(numbers1)
	fmt.Println("Trié ?", sort.IntsAreSorted(numbers1))

	sort.Ints(numbers1)
	fmt.Println(numbers1)
	fmt.Println("Trié ?", sort.IntsAreSorted(numbers1))

	numbers1 = []int{2, 5, 1, 7, 3, 8, 4, 7, 6, 4}
	fmt.Println(numbers1)
	fmt.Println("Trié ?", sort.IntsAreSorted(numbers1))
	sort.Slice(numbers1, func(i, j int) bool {
		return numbers1[i] < numbers1[j]
	})
	fmt.Println(numbers1)
	fmt.Println("Trié ?", sort.IntsAreSorted(numbers1))

	numbers1 = []int{2, 5, 1, 7, 3, 8, 4, 7, 6, 4}
	fmt.Println(numbers1)
	fmt.Println("Trié ?", sort.IntsAreSorted(numbers1))
	// trier en ordre décroissant
	sort.Slice(numbers1, func(i, j int) bool {
		return numbers1[i] > numbers1[j]
	})
	fmt.Println(numbers1)
	fmt.Println("Trié ?", sort.IntsAreSorted(numbers1))

	// trier une tranche de struct Person
	people := []Person{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Alice", Age: 35},
	}
	fmt.Println(people)
	// trier par âge croissant
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
	// trier par âge décroissant
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

	// trier par nom croissant
	sort.Slice(people, personNameCompare(people))
	fmt.Println(people)
	// trier par nom décroissant
	sort.Slice(people, personNameCompareReverse(people))
	fmt.Println(people)

}

func personNameCompareReverse(people []Person) func(i int, j int) bool {
	return func(i, j int) bool {
		return people[i].Name > people[j].Name
	}
}

func personNameCompare(people []Person) func(i int, j int) bool {
	return func(i, j int) bool {
		return people[i].Name < people[j].Name
	}
}

func section4b() {
	people := []Person{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Alice", Age: 35},
	}
	fmt.Println(people)

	// filtrer une tranche de struct Person
	// filtre : plus de 30 ans
	filteredPeople := make([]Person, 0)
	for _, person := range people {
		if person.Age > 30 {
			filteredPeople = append(filteredPeople, person)
		}
	}
	fmt.Println(filteredPeople)

	// filtrer une tranche de struct Person
	// filtre : nom commençant par "J"
	filteredPeople = make([]Person, 0)
	for _, person := range people {
		if person.Name[0] == 'J' {
			filteredPeople = append(filteredPeople, person)
		}
	}
	fmt.Println(filteredPeople)

	// filtrer une tranche de struct Person
	// filtre : nom contenant plus de 5 caractères
	filteredPeople = make([]Person, 0)
	for _, person := range people {
		if len(person.Name) > 5 {
			filteredPeople = append(filteredPeople, person)
		}
	}
	fmt.Println(filteredPeople)
}

func section4c() {
	people := []Person{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Alice", Age: 35},
	}
	fmt.Println(people)

	filteredPeople := filterPeople(people, func(person Person) bool {
		return person.Age > 30
	})
	fmt.Println(filteredPeople)

	filteredPeople = filterPeople(people, func(person Person) bool {
		return person.Name[0] == 'J'
	})
	fmt.Println(filteredPeople)

	filteredPeople = filterPeople(people, func(person Person) bool {
		return len(person.Name) > 5
	})
	fmt.Println(filteredPeople)
}

type FilterFunc func(person Person) bool

func filterPeople(people []Person, filter FilterFunc) []Person {
	filteredPeople := make([]Person, 0)
	for _, person := range people {
		if filter(person) {
			filteredPeople = append(filteredPeople, person)
		}
	}
	return filteredPeople
}

func section4d() {
	people := []Person{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Alice", Age: 35},
	}
	fmt.Println(people)

	filteredPeople := filterPeople(people, peopleFilterMinAge(31))
	fmt.Println(filteredPeople)

}

func peopleFilterMinAge(minAge int) FilterFunc {
	return func(person Person) bool {
		return person.Age >= minAge
	}
}
