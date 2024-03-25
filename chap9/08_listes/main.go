package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	a := app.New()
	w := a.NewWindow("Exemple 08: listes")

	tabs := makeTabs(w)
	w.SetContent(tabs)

	w.Resize(fyne.NewSize(500, 300))

	w.ShowAndRun()
}

func makeTabs(win fyne.Window) *container.AppTabs {
	numbers := makeNumbersList()
	people := makePersonList(win)
	tabs := container.NewAppTabs(
		container.NewTabItem("Numbers", numbers),
		container.NewTabItem("People", people),
	)
	return tabs
}

func makeNumbersList() *container.Scroll {
	numbers := []int{3, 5, 2, 6, 7, 0, 2, -4, 1, 20, -4, 6}
	list := widget.NewList(
		func() int {
			return len(numbers)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Item")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fmt.Sprintf("Item %d : %d", id, numbers[id]))
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		numbers[id]++
		list.RefreshItem(id)
	}
	return container.NewVScroll(list)
}

func makePersonList(win fyne.Window) *container.Scroll {
	people := []Person{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Alice", Age: 35},
	}
	list := widget.NewList(
		func() int {
			return len(people)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Person")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			person := people[id]
			item.(*widget.Label).SetText(fmt.Sprintf("Name: %s, Age: %d", person.Name, person.Age))
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		out := fmt.Sprintf("Personne sélectionnée : %s", people[id].Name)
		dialog.ShowInformation(fmt.Sprintf("Personne %d", id), out, win)
	}
	return container.NewVScroll(list)
}
