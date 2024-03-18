package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Allo")
	w.SetContent(widget.NewLabel("Allo"))
	w.SetMaster()
	w.Show()

	w2 := a.NewWindow("Plus grande")
	button := widget.NewButton("Ouvrir", func() {
		w3 := a.NewWindow("Trois")
		w3.SetContent(widget.NewLabel("Trois"))
		w3.Resize(fyne.NewSize(100, 100))
		w3.Show()
	})
	w2.SetContent(button)
	w2.Resize(fyne.NewSize(200, 200))
	w2.Show()

	a.Run()
}
