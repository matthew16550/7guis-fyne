package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"time"
)

// Not using fyne.Animation because Animation.Duration is only used at the tick when an animation reverses or repeats so it does not help here

// TODO slider width

const frameRate = 25
const maxDurationSeconds = 120

func main() {
	window := app.New().NewWindow("Timer")

	start := time.Now()

	elapsed := widget.NewProgressBar()
	elapsed.TextFormatter = func() string { return fmt.Sprintf("%.1fs", elapsed.Value) }

	duration := widget.NewSlider(0, maxDurationSeconds)
	duration.OnChanged = func(f float64) { elapsed.Max = f }
	duration.SetValue(duration.Max / 2)

	reset := widget.NewButton("Reset", func() {
		start = time.Now()
	})

	window.SetContent(
		container.NewCenter(
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Elapsed Time"), elapsed,
				widget.NewLabel("Duration"), duration,
				widget.NewLabel(""), reset,
			),
		),
	)

	ticker := time.NewTicker(time.Second / frameRate)

	go func() {
		for now := range ticker.C {
			elapsed.SetValue(now.Sub(start).Seconds())
		}
	}()

	window.ShowAndRun()
}
