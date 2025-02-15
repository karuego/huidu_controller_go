package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	xlay "fyne.io/x/fyne/layout"
)

func window_home(w fyne.Window) {
	selection_device := container.NewHBox(
		widget.NewLabel("Device terpilih:"),
		widget.NewLabelWithData(selected_device),
		/*widget.NewSelect([]string{"Top", "Bottom", "Leading", "Trailing"}, func(s string) {
			log.Printf("Device terpilih: %s", s)
		}),*/
		layout.NewSpacer(),
		window_scan(&myApp, w),
	)

	borderLine := canvas.NewLine(color.Gray{0x80})
	borderLine.StrokeWidth = 2
	borderLine.Resize(fyne.NewSize(selection_device.MinSize().Width, 2))
	borderLine.Move(fyne.NewPos(0, selection_device.MinSize().Height))

	btn_kirim_teks := widget.NewButton("Kirim Teks", func() {
		window_send_text(w)
	})
	btn_kirim_teks.Importance = widget.HighImportance

	w.SetContent(xlay.NewResponsiveLayout(container.NewVBox(
		selection_device,
		borderLine,
		widget.NewSeparator(),
		btn_kirim_teks,
	)))
}
