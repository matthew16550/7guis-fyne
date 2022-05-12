package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"sync"
)

func celsius2fahrenheit(c float64) float64 {
	return c*(9.0/5.0) + 32.0
}

func fahrenheit2celsius(f float64) float64 {
	return (f - 32.0) * (5.0 / 9.0)
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 32)
}

func formatFloat(f float64) string {
	return fmt.Sprintf("%.1f", f)
}

type widgets struct {
	celsius    *widget.Entry
	fahrenheit *widget.Entry
}

func makeUI() widgets {
	celsius := widget.NewEntry()

	fahrenheit := widget.NewEntry()

	var lock sync.Mutex // prevents Entry.SetText() causing a loop

	celsius.OnChanged = func(s string) {
		if !lock.TryLock() {
			return
		}
		defer lock.Unlock()

		c, err := parseFloat(s)
		if err != nil {
			return
		}

		fahrenheit.SetText(formatFloat(celsius2fahrenheit(c)))
	}

	fahrenheit.OnChanged = func(s string) {
		if !lock.TryLock() {
			return
		}
		defer lock.Unlock()

		f, err := parseFloat(s)
		if err != nil {
			return
		}

		celsius.SetText(formatFloat(fahrenheit2celsius(f)))
	}

	return widgets{celsius, fahrenheit}
}

func main() {
	ui := makeUI()

	window := app.New().NewWindow("Temperature Converter")

	// TODO not clear how to set min width of entry boxes
	window.SetContent(
		container.NewCenter(
			container.NewHBox(
				ui.celsius,
				widget.NewLabel("Celsius ="),
				ui.fahrenheit,
				widget.NewLabel("Fahrenheit"),
			),
		),
	)

	window.ShowAndRun()
}
