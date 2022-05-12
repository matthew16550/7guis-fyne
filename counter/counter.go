package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// TODO not clear how to test bindings, think we need to call binding.waitForItems() at start of the test

func main() {
	count := binding.NewInt()

	label := widget.NewLabelWithData(binding.IntToString(count))

	button := widget.NewButton("Count", func() {
		c, _ := count.Get()
		_ = count.Set(c + 1)
	})

	window := app.New().NewWindow("Counter")

	window.SetContent(
		container.NewCenter(
			container.New(layout.NewFormLayout(), label, button),
		),
	)

	window.Resize(fyne.NewSize(200, 50)) // kludge to make the whole window title visible
	window.ShowAndRun()
}
