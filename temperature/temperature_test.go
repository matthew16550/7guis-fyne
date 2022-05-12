package main

import (
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemperatureConverter(t *testing.T) {
	ui := makeUI()

	assert.Equal(t, "", ui.celsius.Text, "C initially empty")
	assert.Equal(t, "", ui.fahrenheit.Text, "F initially empty")

	test.Type(ui.celsius, "5.0")
	assert.Equal(t, "41.0", ui.fahrenheit.Text, "5.0 C to 41.0 F")

	test.Type(ui.celsius, "foo")
	assert.Equal(t, "41.0", ui.fahrenheit.Text, "Invalid C does not change F")

	ui.celsius.SetText("")
	ui.fahrenheit.SetText("")

	test.Type(ui.fahrenheit, "45.0")
	assert.Equal(t, "7.2", ui.celsius.Text, "45.0 F to 7.2 C")

	test.Type(ui.fahrenheit, "foo")
	assert.Equal(t, "7.2", ui.celsius.Text, "Invalid F does not change C")
}
