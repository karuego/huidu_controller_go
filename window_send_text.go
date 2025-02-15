package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func window_send_text(w fyne.Window) {
	//w := myApp.NewWindow("Kirim Teks")
	//w.Resize(fyne.NewSize(360, 240))
	// defer w.Close()

	input := widget.NewEntry()
	input.SetPlaceHolder("Masukkan teks...")

	btn_kirim := widget.NewButton("Kirim", func() {
		if len(input.Text) < 1 {
			return
		}
		log.Println("Content was:", input.Text)
	})
	btn_kirim.Importance = widget.HighImportance

	input.OnSubmitted = func(teks string) {
		btn_kirim.OnTapped()
	}

	content := container.NewVBox(
		tombol_kembali(w),
		input,
		btn_kirim,
	)

	w.SetContent(content)
	//go func() { w.Show() }()
}
