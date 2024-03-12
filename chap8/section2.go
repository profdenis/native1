package main

import (
	"fmt"
	"sync"
	"time"
)

func countdown(n int, id int) {
	for i := n; i >= 0; i-- {
		fmt.Printf("goroutine %d : %d\n", id, i)
		time.Sleep(1 * time.Second)
	}
}

func section2a() {
	fmt.Println("Simple")
	for i := 1; i <= 4; i++ {
		go countdown(5, i) // Lancer la goroutine
	}
	time.Sleep(6 * time.Second) // Laisser du temps pour l'exécution des goroutines
}

// Nous ajoutons un canal done de type bool
func countdownChan(n int, id int, done chan<- bool) {
	for i := n; i >= 0; i-- {
		fmt.Printf("goroutine %d : %d\n", id, i)
		time.Sleep(1 * time.Second)
	}
	done <- true // Envoie true sur le canal quand la goroutine a terminé
}

func section2b() {
	fmt.Println("Chan")
	done := make(chan bool, 4) // Crée un canal de booléens avec une capacité de 4

	for i := 1; i <= 4; i++ {
		go countdownChan(5, i, done) // démarre une goroutine et lui passe le canal done
	}

	for i := 1; i <= 4; i++ {
		<-done // Attend un vrai de chaque goroutine
	}
}

func countdownWaitGroup(n int, id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signale au WaitGroup que cette goroutine est terminée lorsque la fonction se termine
	for i := n; i >= 0; i-- {
		fmt.Printf("goroutine %d : %d\n", id, i)
		time.Sleep(1 * time.Second)
	}
}

func section2c() {
	fmt.Println("WaitGroup")
	var wg sync.WaitGroup // Crée un WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)                        // Augmente le compteur du WaitGroup de 1
		go countdownWaitGroup(5, i, &wg) // Démarre une goroutine avec le WaitGroup
	}

	wg.Wait() // Attend que toutes les goroutines signalent qu'elles sont terminées
}
