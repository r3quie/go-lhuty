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
	dnes := time.Now()

	title := widget.NewLabel("Kalkulačka lhůt")
	test := widget.NewLabel("Zadejte datum a počet dní")

	zacatek := widget.NewEntry()
	zacatek.SetPlaceHolder(f("%d.%d.%d", dnes.Day(), dnes.Month(), dnes.Year()))

	dnu := widget.NewEntry()
	dnu.SetPlaceHolder("8")

	vysledek := widget.NewLabel("")

	w.SetContent(container.NewVBox(
		title,
		test,
		zacatek,
		dnu,
		vysledek,
		widget.NewButton("Spočítat", func() {
			vysledek.SetText(doruceni(string_to_time(zacatek.Text), string_to_int(dnu.Text)))
		}),
	))

	w.ShowAndRun()
}
