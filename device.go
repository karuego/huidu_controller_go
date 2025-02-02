package main

import (
	"time"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/data/binding"
)

var selected_device binding.String = binding.NewString()

// TODO: gunakan binding
func searchDevice(list_data *[]string, list *widget.List, fn func()) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)

		(*list_data) = append(*list_data, "Test Item " + strconv.Itoa(i))
		list.Refresh()
	}

	fn()
}

type resizeRefreshCountingLabel struct {
	widget.Label
	resizeCount  int
	refreshCount int
}

func newResizeRefreshCountingLabel(text string) *resizeRefreshCountingLabel {
	r := &resizeRefreshCountingLabel{}
	r.Text = text
	r.ExtendBaseWidget(r)
	return r
}

func (r *resizeRefreshCountingLabel) Refresh() {
	r.refreshCount++
	r.Label.Refresh()
}

func (r *resizeRefreshCountingLabel) Resize(s fyne.Size) {
	r.resizeCount++
	r.Label.Resize(s)
}
