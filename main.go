package main

import (
	// "log"
	// "time"
	// "strconv"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/data/binding"

	xlay "fyne.io/x/fyne/layout"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Skripsi")
	w.Resize(fyne.NewSize(480, 480))
	// defer w.Close()

	/*btn_kirim_teks := widget.NewButton("Kirim Teks", func() {
		w := a.NewWindow("Kirim Teks")
		defer w.Close()
		input := widget.NewEntry()
		input.SetPlaceHolder("Masukkan teks...")

		content := container.NewVBox(input, widget.NewButton("Kirim", func() {
			log.Println("Content was:", input.Text)
		}))
		w.SetContent(content)
		go func() { w.Show() }()
	})*/

	selected_device.Set("------")

	selection_device := container.NewHBox(
		widget.NewLabel("Device terpilih:"),
		widget.NewLabelWithData(selected_device),
		/*widget.NewSelect([]string{"Top", "Bottom", "Leading", "Trailing"}, func(s string) {
			log.Printf("Device terpilih: %s", s)
		}),*/
		layout.NewSpacer(),
		window_search_device(&a),
	)

	borderLine := canvas.NewLine(color.Gray{0x80})
	borderLine.StrokeWidth = 2
	borderLine.Resize(fyne.NewSize(selection_device.MinSize().Width, 2))
	borderLine.Move(fyne.NewPos(0, selection_device.MinSize().Height))

	w.SetContent(xlay.NewResponsiveLayout(container.NewVBox(
		selection_device,
		borderLine,
	)))
	w.ShowAndRun()
	//w.Show()
	//a.Run()
}


func window_search_device(a *fyne.App) fyne.CanvasObject {
	btn := widget.NewButton("Pindai device", func() {
		list_data := make([]string, 0, 32)
		selected_device.Set("-----")

		list := widget.NewList(
			func() int {
				return len(list_data)
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("Template object")
			},
			func(id widget.ListItemID, item fyne.CanvasObject) {
				item.(*widget.Label).SetText(list_data[id])
			},
		)
		list.OnSelected = func(id widget.ListItemID) {
			selected_device.Set(list_data[id])
		}
		list.OnUnselected = func(id widget.ListItemID) {
			selected_device.Set("-----")
		}
		// list.Select(125)
		// list.SetItemHeight(5, 50)
		// list.SetItemHeight(6, 50)
		// list.Resize(fyne.NewSize(200, 150))

		w_list_device := (*a).NewWindow("Pilih controller")
		defer w_list_device.Close()
		w_list_device.Resize(fyne.NewSize(550, 500))

		progress := widget.NewProgressBarInfinite()
		progress.Start()

		fetch := func() {
			go searchDevice(&list_data, list, func() {
				progress.Stop()
				progress.Hide()
			})
		}
		fetch()

		btn_ok := widget.NewButton("   Ok   ", func() {
			w_list_device.Close()
		})
		btn_ok.Importance = widget.HighImportance

		btn_refresh := widget.NewButton("Refresh", func() {
			list_data = []string{}
			fetch()
			progress.Start()
			progress.Show()
		})
		// btn_refresh.Importance = widget.MImportance

		selection_device := container.NewBorder(
			nil, nil,
			widget.NewLabel("Device terpilih: "),
			// btn_ok,
			container.NewBorder(
				nil, nil,
				btn_refresh,
				layout.NewSpacer(),
				btn_ok,
			),
			// layout.NewSpacer(),
			// selected_device,
			widget.NewLabelWithData(selected_device),
		)

		borderLine := canvas.NewLine(color.Gray{0x80})
		borderLine.StrokeWidth = 2
		borderLine.Resize(fyne.NewSize(selection_device.MinSize().Width, 2))
		borderLine.Move(fyne.NewPos(0, selection_device.MinSize().Height))

		w_list_device.SetContent(
			container.NewVScroll(
				container.NewBorder(
					selection_device,
					widget.NewSeparator(),
					nil, nil,
					container.NewBorder(
						borderLine,
						nil, nil, nil,
					),
					// list,
					container.NewBorder(
						nil,
						progress,
						nil, nil,
						list,
					),
				),
			),
		)

		w_list_device.Show()
	})

	btn.Importance = widget.HighImportance
	return btn
}
