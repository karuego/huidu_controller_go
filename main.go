package main

import (
	// "time"
	// "strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var myApp fyne.App = app.New()

var isActive map[string]bool

func main() {
	//myApp.Settings().SetTheme()
	w := myApp.NewWindow("Skripsi")
	w.Resize(fyne.NewSize(480, 480))
	// defer w.Close()

	isActive = map[string]bool{
		"home":     true,
		"scan":     false,
		"send_txt": false,
		"send_img": false,
		"send_vid": false,
		"settings": false,
	}

	selected_device.Set("------")

	window_home(w)
	w.ShowAndRun()
	//w.Show()
	//a.Run()
}
