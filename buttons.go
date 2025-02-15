package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func tombol_kembali(w fyne.Window) fyne.Widget {
	btn := widget.NewButton("Kembali", func() {
		window_home(w)
	})
	btn.Importance = widget.HighImportance

	return btn
}
