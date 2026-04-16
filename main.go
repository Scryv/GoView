package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"io/ioutil"
	"log"
	"os"
)

func folderSpy() (data []string, path string) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if len(name) < 4 {
			log.Println(name + " too short")
			continue
		}

		switch name[len(name)-4:] {
		case ".jpg":
			data = append(data, name)
		case ".png":
			data = append(data, name)
		default:
			log.Println(name + "Not an image")
		}
	}
	return
}

func showImage(w fyne.Window, path string) {
	image := canvas.NewImageFromFile(path)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)
}

func main() {
	var current int
	a := app.New()
	w := a.NewWindow("GoView")

	fileData, filePath := folderSpy()
	log.Println(fileData, filePath)

	path := filePath + "/" + fileData[current]

	image := canvas.NewImageFromFile("images.jpg")
	image.FillMode = canvas.ImageFillOriginal

	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	ctrlAltTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl | fyne.KeyModifierAlt}

	w.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		if current < len(fileData)-1 {
			current++
		}
		path := filePath + "/" + fileData[current]
		fmt.Println(path)
		showImage(w, path)
	})
	w.Canvas().AddShortcut(ctrlAltTab, func(shortcut fyne.Shortcut) {
		if current > 0 {
			current--
		}
		path := filePath + "/" + fileData[current]
		showImage(w, path)
	})

	showImage(w, path)

	w.ShowAndRun()

}
