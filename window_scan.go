package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func window_scan(myApp *fyne.App, w fyne.Window) fyne.CanvasObject {
	btn := widget.NewButton("Pindai device", func() {
		// list_data := make([]string, 0, 32)
		list_data := binding.NewStringList()
		selected_device.Set("-----")

		list := widget.NewListWithData(
			list_data,
			func() fyne.CanvasObject {
				return widget.NewLabel("Template object")
			},
			func(i binding.DataItem, o fyne.CanvasObject) {
				o.(*widget.Label).Bind(i.(binding.String))
			},
		)
		list.OnSelected = func(id widget.ListItemID) {
			data, err := list_data.GetValue(id)
			if err != nil {
				log.Printf("Error: %s", err)
			}
			selected_device.Set(data)
		}
		list.OnUnselected = func(id widget.ListItemID) {
			selected_device.Set("-----")
		}
		// list.Select(125)
		// list.SetItemHeight(5, 50)
		// list.SetItemHeight(6, 50)
		// list.Resize(fyne.NewSize(200, 150))

		w_list_device := (*myApp).NewWindow("Pilih controller")
		defer w_list_device.Close()
		w_list_device.Resize(fyne.NewSize(550, 500))

		progress := widget.NewProgressBarInfinite()
		progress.Start()

		fetch := func() {
			go searchDevice(w_list_device, &list_data, func() {
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
			list.UnselectAll()
			list_data.Set([]string{})
			selected_device.Set("-----")
			fetch()
			progress.Start()
			progress.Show()
		})
		// btn_refresh.Importance = widget.MImportance

		selection_device := container.NewBorder(
			tombol_kembali(w),
			nil,
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
