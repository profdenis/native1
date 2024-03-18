package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI() *fyne.Container {
	out := widget.NewLabel("Texte : ")
	in := widget.NewEntry()
	in.OnChanged = func(content string) {
		out.SetText(fmt.Sprintf("Texte : %s", content))
	}
	//in.Password = true
	return container.NewVBox(out, in)
}

func main() {
	a := app.New()
	w := a.NewWindow("Événement")
	w.SetContent(makeUI())
	w.ShowAndRun()
}

//exercice : faire une fenêtre de connexion avec des entrées pour l'utilisateur et le mot de passe,
//et 2 boutons (annuler et connecter)
//annuler : fermer la fenêtre
//connecter : si le mot de passe est 123456, ouvrir un autre fenêtre contenant le texte "Réussi !" et fermer la fenêtre
//de connexion ; sinon, indiquer une erreur dans la même fenêtre sous le mot de passe
