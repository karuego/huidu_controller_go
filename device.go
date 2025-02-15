package main

import (
	// "time"
	// "strconv"
	"fmt"
	"net"
	"sync"
	"errors"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/data/binding"
)

var selected_device binding.String = binding.NewString()

// TODO: gunakan binding
func searchDevice(parent fyne.Window, list_data *binding.StringList, fn func()) {
	var wg sync.WaitGroup
	deviceCh := make(chan string)
	errCh := make(chan error)
	go searchDeviceAsk(deviceCh, errCh)

	fmt.Println("satu")
	select {
	case err_ := <-errCh:
		fmt.Printf("err type: %T\n", err_)
		switch err := err_.(type) {
		case (*net.OpError):
			if errors.Is(err, syscall.ENETUNREACH) {
				fmt.Println("Detail error: ", err)
				dialog.ShowError(errors.New("Error: Jaringan tidak dapat dijangkau"), parent)
			}
		}
	default:

	}

	fmt.Println("dua")
	devices := make(map[string]struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for device := range deviceCh {
			if _, exists := devices[device]; exists {
				continue
			}

			devices[device] = struct{}{}
			(*list_data).Append(device)
		}
	}()

	wg.Wait()

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
