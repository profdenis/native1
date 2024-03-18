package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func makeHBoxContainer() *fyne.Container {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")

	return container.New(layout.NewHBoxLayout(), label1, label2, label3)
}

func makeVBoxContainer() *fyne.Container {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")

	return container.NewVBox(label1, label2, label3)
}

func makeCenterContainer() *fyne.Container {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")

	// normalement préférable de donner seulement un élément à NewCenter
	return container.NewCenter(label1, label2, label3)
}

func makeCenterVBoxContainer() *fyne.Container {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")

	return container.NewCenter(container.NewVBox(label1, label2, label3))
}

func makeBorderContainer() *fyne.Container {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")
	label4 := widget.NewLabel("Quatre")
	label5 := widget.NewLabel("Cinq")

	return container.NewBorder(label1, label2, label3, label4, label5)
}

func makeBorderCenterContainer() *fyne.Container {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")
	label4 := widget.NewLabel("Quatre")
	label5 := widget.NewLabel("Cinq")

	return container.NewBorder(label1, label2, label3, label4, container.NewCenter(label5))
}

func makeGridWrapContainer() *fyne.Container {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")
	label4 := widget.NewLabel("Quatre")
	label5 := widget.NewLabel("Cinq")
	label6 := widget.NewLabel("Six")

	return container.NewGridWrap(fyne.NewSize(100, 50), label1, label2, label3, label4, label5, label6)
}

func makeGridWithRowsContainer() *fyne.Container {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")
	label4 := widget.NewLabel("Quatre")
	label5 := widget.NewLabel("Cinq")
	label6 := widget.NewLabel("Six")

	return container.NewGridWithRows(3, label1, label2, label3, label4, label5, label6)
}

func makeGridWithColumnsContainer() *fyne.Container {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")
	label4 := widget.NewLabel("Quatre")
	label5 := widget.NewLabel("Cinq")
	label6 := widget.NewLabel("Six")

	return container.NewGridWithColumns(3, label1, label2, label3, label4, label5, label6)
}

func makeHSplitContainer() *container.Split {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")

	return container.NewHSplit(label1, label2)
}

func makeVSplitContainer() *container.Split {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")

	return container.NewVSplit(label1, label2)
}

func makeVScrollContainer() *container.Scroll {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")
	label4 := widget.NewLabel("Quatre")
	label5 := widget.NewLabel("Cinq")
	label6 := widget.NewLabel("Six")

	return container.NewVScroll(container.NewVBox(label1, label2, label3, label4, label5, label6))
}

func makeHScrollContainer() *container.Scroll {
	label1 := widget.NewLabel("Un")
	label2 := widget.NewLabel("Deux")
	label3 := widget.NewLabel("Trois")
	label4 := widget.NewLabel("Quatre")
	label5 := widget.NewLabel("Cinq")
	label6 := widget.NewLabel("Six")

	return container.NewHScroll(container.NewHBox(label1, label2, label3, label4, label5, label6))
}

func main() {
	a := app.New()
	w := a.NewWindow("Container")
	//w.SetContent(makeHBoxContainer())
	//w.SetContent(makeVBoxContainer())
	//w.SetContent(makeCenterContainer())
	//w.SetContent(makeCenterVBoxContainer())
	//w.SetContent(makeBorderContainer())
	//w.SetContent(makeBorderCenterContainer())
	//w.SetContent(makeGridWrapContainer())
	//w.SetContent(makeGridWithRowsContainer())
	//w.SetContent(makeGridWithColumnsContainer())
	//w.SetContent(makeHSplitContainer())
	//w.SetContent(makeVSplitContainer())
	//w.SetContent(makeVScrollContainer())
	w.SetContent(makeHScrollContainer())
	w.ShowAndRun()
}
