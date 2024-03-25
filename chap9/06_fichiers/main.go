package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"io"
	"log"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("Exemple 06 : Fichiers")

	src, dst := makeSrcDstEntries()

	split := container.NewHSplit(src, dst)
	w.SetContent(container.NewBorder(
		container.NewCenter(
			container.NewHBox(
				widget.NewButton("Ouvrir", getOpenButtonFunc(w, src, dst)),
				//TODO: "refactor extract method", sur 2 fonctions anonymes
				widget.NewButton("Enregistrer", func() {
					fd := dialog.NewFileSave(func(closer fyne.URIWriteCloser, err error) {
						if err != nil {
							dialog.ShowError(err, w)
							return
						}
						if closer == nil {
							log.Println("Cancelled")
							return
						}

						file, err := os.Create(closer.URI().Path())
						if err != nil {
							dialog.ShowError(err, w)
							return
						}

						defer file.Close()

						_, err = file.Write([]byte(src.Text))
						if err != nil {
							dialog.ShowError(err, w)
						}
					}, w)
					fd.SetFilter(storage.NewExtensionFileFilter([]string{".md"}))

					listableUri, err := getCurrentDirURI()
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
					fd.SetLocation(listableUri)
					fd.SetFileName("example.md")

					fd.Show()
				}),
			),
		),
		nil,
		nil,
		nil,
		split,
	))

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func getOpenButtonFunc(w fyne.Window, src *widget.Entry, dst *widget.RichText) func() {
	return func() {
		fd := dialog.NewFileOpen(getOpenFileDialogFunc(w, src, dst), w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".md"}))

		listableUri, err := getCurrentDirURI()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		fd.SetLocation(listableUri)

		fd.Show()
	}
}

func getOpenFileDialogFunc(w fyne.Window, src *widget.Entry, dst *widget.RichText) func(reader fyne.URIReadCloser, err error) {
	return func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if reader == nil {
			log.Println("Cancelled")
			return
		}
		data, err := io.ReadAll(reader)
		if err != nil {
			fmt.Println(err)
		}
		src.SetText(string(data))
		dst.ParseMarkdown(string(data))
	}
}

func makeSrcDstEntries() (*widget.Entry, *widget.RichText) {
	src := widget.NewEntry()
	src.MultiLine = true
	src.SetMinRowsVisible(10)
	src.Scroll = container.ScrollVerticalOnly
	src.Wrapping = fyne.TextWrapWord
	dst := widget.NewRichTextFromMarkdown("")
	dst.Scroll = container.ScrollVerticalOnly
	dst.Wrapping = fyne.TextWrapWord

	src.OnChanged = func(content string) {
		dst.ParseMarkdown(content)
	}
	return src, dst
}

func getCurrentDirURI() (fyne.ListableURI, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	uri := storage.NewFileURI(dir)
	listableUri, err := storage.ListerForURI(uri)
	if err != nil {
		return nil, err
	}

	return listableUri, nil
}
