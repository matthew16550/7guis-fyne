package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"time"
)

// TODO not simple to make error backgrounds red
// TODO resize so combo box all visible

const (
	oneWayFlight = "one-way flight"
	returnFlight = "return flight"
)

const minYear = 2000
const maxYear = 2099

type inputs struct {
	flightType string
	startDate  string
	returnDate string
}

func (b *inputs) isValid() bool {
	if b.flightType == oneWayFlight {
		_, err := parseDate(b.startDate)
		return err == nil
	} else {
		startDate, startErr := parseDate(b.startDate)
		returnDate, returnErr := parseDate(b.returnDate)
		return startErr == nil && returnErr == nil && !returnDate.Before(startDate)
	}
}

func parseDate(s string) (time.Time, error) {
	date, err := time.Parse("02.01.2006", s)
	if err != nil {
		return date, err
	}
	if date.Year() < minYear || date.Year() > maxYear {
		err := errors.New("bad year")
		return date, err
	}
	return date, nil
}

func main() {
	window := app.New().NewWindow("Book Flight")

	flightType := widget.NewSelect([]string{oneWayFlight, returnFlight}, nil)

	startDate := widget.NewEntry()
	startDate.Text = "27.03.2014"

	returnDate := widget.NewEntry()
	returnDate.Text = "27.03.2014"

	toInputs := func() *inputs {
		return &inputs{
			flightType: flightType.Selected,
			startDate:  startDate.Text,
			returnDate: returnDate.Text,
		}
	}

	bookButton := widget.NewButton("Book", func() {
		b := toInputs()
		if !b.isValid() {
			return // should never happen
		}

		var message string

		if b.flightType == oneWayFlight {
			message = fmt.Sprintf("You have booked a one-way flight on %s.", b.startDate)
		} else {
			message = fmt.Sprintf("You have booked a return flight starting %s and returning %s.", b.startDate, b.returnDate)
		}

		dialog.ShowInformation("Booking created", message, window)
	})

	update := func(_ string) {
		setWidgetEnabled(returnDate, flightType.Selected == returnFlight)
		setWidgetEnabled(bookButton, toInputs().isValid())
	}

	flightType.OnChanged = update
	startDate.OnChanged = update
	returnDate.OnChanged = update

	flightType.SetSelected(oneWayFlight) // triggers update()

	window.SetContent(
		container.NewCenter(
			container.NewVBox(
				flightType,
				startDate,
				returnDate,
				bookButton,
			),
		),
	)

	window.Resize(fyne.NewSize(800, 10)) // kludge to ensure booking dialog is all visible
	window.ShowAndRun()
}

func setWidgetEnabled(d fyne.Disableable, enabled bool) {
	if enabled {
		d.Enable()
	} else {
		d.Disable()
	}
}
