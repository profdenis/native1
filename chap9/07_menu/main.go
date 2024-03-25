package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"io"
	"log"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("Exemple 07 : Menu et raccourcis")

	filename := widget.NewLabel("")
	src, dst := makeSrcDstEntries()

	//TODO: "refactor extract method" pour construire mainMenu
	// Créer la barre de menus
	openMenuItem := fyne.NewMenuItem("Open", getOpenMenuFunc(w, src, dst, filename))
	openMenuItem.Icon = theme.DocumentIcon()
	openMenuItem.Shortcut = &desktop.CustomShortcut{
		KeyName:  fyne.KeyO,
		Modifier: fyne.KeyModifierShortcutDefault,
	}
	w.Canvas().AddShortcut(openMenuItem.Shortcut, func(shortcut fyne.Shortcut) {
		getOpenMenuFunc(w, src, dst, filename)()
	})

	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			openMenuItem,
			//TODO: "refactor introduce variable", et ajouter un icône et un raccourci
			fyne.NewMenuItem("Save", getSaveMenuFunc(w, src)),
		),
		// ajouter d'autres menus
		//fyne.NewMenu("Edit", openMenuItem),
	)

	// placer le menu principal dans la fenêtre
	w.SetMainMenu(mainMenu)

	split := container.NewHSplit(src, dst)
	w.SetContent(container.NewBorder(filename, nil, nil, nil, split))

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func getSaveMenuFunc(w fyne.Window, src *widget.Entry) func() {
	return func() {
		fd := dialog.NewFileSave(getFileSaveDialogFunc(w, src), w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".md"}))

		//TODO: s'il n'y a pas de fichier ouvert, alors utiliser le dossier courant dans SetLocation
		// s'il y a un fichier ouvert, utiliser le dossier de ce fichier dans SetLocation, et le nom du fichier dans
		// SetFileName
		listableUri, err := getCurrentDirURI()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		fd.SetLocation(listableUri)
		//TODO: "refactor extract method"
		//code partagé avec le dialogue d'ouverture de fichier (sauf pour le nom du fichier)

		fd.SetFileName("example.md")

		fd.Show()
	}
}

func getFileSaveDialogFunc(w fyne.Window, src *widget.Entry) func(closer fyne.URIWriteCloser, err error) {
	return func(closer fyne.URIWriteCloser, err error) {
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
	}
}

func getOpenMenuFunc(w fyne.Window, src *widget.Entry, dst *widget.RichText, filename *widget.Label) func() {
	return func() {
		fd := dialog.NewFileOpen(getOpenFileDialogFunc(w, src, dst, filename), w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".md"}))

		//TODO: s'il n'y a pas de fichier ouvert, alors utiliser le dossier courant dans SetLocation
		// s'il y a un fichier ouvert, utiliser le dossier de ce fichier dans SetLocation
		listableUri, err := getCurrentDirURI()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		fd.SetLocation(listableUri)

		fd.Show()
	}
}

func getOpenFileDialogFunc(w fyne.Window, src *widget.Entry, dst *widget.RichText, filename *widget.Label) func(reader fyne.URIReadCloser, err error) {
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
		filename.SetText(reader.URI().Path())
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
