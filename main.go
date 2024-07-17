package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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

	zacatekstr := binding.NewString()
	zacatekstr.Set(f("%d.%d.%d", dnes.Day(), dnes.Month(), dnes.Year()))
	dnustr := binding.NewString()
	sadasd := binding.FloatToStringWithFormat(0, "%.0f")
	int1 := binding.NewInt()
	int1.Set(0)
	asdadsdd := binding.IntToString(int1)
	dnustr.Set("8")
	vysledekstr := binding.NewString()

	title := widget.NewLabel("Kalkulačka lhůt")
	test := widget.NewLabel("Zadejte datum a počet dní")

	zacatek := widget.NewEntryWithData(zacatekstr)
	dnu := widget.NewEntryWithData(dnustr)

	//workv, _ := zacatekstr.Get()
	//workdnu, _ := dnustr.Get()

	b, _ := dnustr.Get()
	vysledekstr.Set(doruceni(string_to_time(a), string_to_int(b)))

	vysledek := widget.NewLabelWithData(vysledekstr)

	w.SetContent(container.NewVBox(
		title,
		test,
		zacatek,
		dnu,
		vysledek,
	))

	w.ShowAndRun()
}
