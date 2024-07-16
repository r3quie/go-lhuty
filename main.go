package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Kalkulačka lhůt")

	title := widget.NewLabel("Kalkulačka lhůt")
	test := widget.NewLabel("12.4.2024")
	w.SetContent(container.NewVBox(
		title,
		test,
		widget.NewButton("Spočítat", func() {
			test.SetText(doruceni(time.Now(), 5))
		}),
	))

	w.ShowAndRun()
}
