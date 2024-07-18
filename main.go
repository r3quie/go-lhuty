package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*
	type base struct {
		listeners sync.Map // map[DataListener]bool

		lock sync.RWMutex
	}

	type Doruceni interface {
		binding.DataItem
		Get() (string, error)
		Set(string) error
	}

	type stringToDoruceni struct {
		base
		format string
		from   string
	}

	func StringToDoruceni(str String) binding.String {
		v := &stringToDoruceni{from: str}
		str.binding.AddListener(v)
		return v
	}
*/

func main() {
	a := app.New()
	w := a.NewWindow("Kalkulačka lhůt")
	dnes := time.Now()

	title := widget.NewLabel("Kalkulačka lhůt")
	test := widget.NewLabel("Zadejte datum a počet dní")

	zacatek := widget.NewEntry()
	zacatek.PlaceHolder = strformat(dnes)
	dnu := widget.NewEntry()
	dnu.PlaceHolder = "8"

	//workv, _ := zacatekstr.Get()
	//workdnu, _ := dnustr.Get()

	vysledek := widget.NewLabel(doruceni(dnes, 8))

	w.SetContent(container.NewVBox(
		title,
		test,
		zacatek,
		dnu,
		vysledek,
	))

	dnu.OnChanged = func(s string) {
		vysledek.SetText(doruceni(string_to_time(zacatek.Text), string_to_int(s)))
	}
	zacatek.OnChanged = func(ss string) {
		vysledek.SetText(doruceni(string_to_time(ss), string_to_int(dnu.Text)))
	}

	w.ShowAndRun()
}
