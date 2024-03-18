package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Allo")

	allo := widget.NewLabel("Allo")
	w.SetContent(allo)
	//w.SetContent(container.NewVBox(
	//	allo,
	//	widget.NewButton("Salut !", func() {
	//		allo.SetText("Bienvenue 😀")
	//	}),
	//))
	w.ShowAndRun()
}
