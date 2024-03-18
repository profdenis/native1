package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTimeLabel(clock)

	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTimeLabel(clock)
		}
	}()
	w.ShowAndRun()
}

func updateTimeLabel(clock *widget.Label) {
	formatted := time.Now().Format("L'heure est 03:04:05")
	clock.SetText(formatted)
}
