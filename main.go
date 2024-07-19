package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Kalkulačka lhůt")
	dnes := time.Now()

	title := widget.NewLabel("Kalkulačka lhůt")
	labeldat := widget.NewLabel("Zadejte datum")
	labeldelta := widget.NewLabel("Zadejte počet dnů")

	zacatek := widget.NewEntry()
	zacatek.PlaceHolder = strformat(dnes)
	dnu := widget.NewEntry()
	dnu.PlaceHolder = "8"

	vysledek := widget.NewLabel(doruceni(dnes, 8))

	w.SetContent(container.NewVBox(
		title,
		labeldat,
		zacatek,
		labeldelta,
		dnu,
		vysledek,
	))

	w.Resize(fyne.NewSize(600, 200))

	dnu.OnChanged = func(s string) {
		vysledek.SetText(doruceni(string_to_time(zacatek.Text), string_to_int(s)))
	}
	zacatek.OnChanged = func(ss string) {
		vysledek.SetText(doruceni(string_to_time(ss), string_to_int(dnu.Text)))
	}

	w.ShowAndRun()
}
