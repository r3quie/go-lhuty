package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Kalkulačka lhůt")
	dnes := time.Now()

	zacatekstr := binding.NewString()
	zacatekstr.Set(f("%d.%d.%d", dnes.Day(), dnes.Month(), dnes.Year()))
	dnustr := binding.NewString()
	dnustr.Set("8")
	vysledekstr := binding.NewString()

	title := widget.NewLabel("Kalkulačka lhůt")
	test := widget.NewLabel("Zadejte datum a počet dní")

	zacatek := widget.NewEntryWithData(zacatekstr)

	dnu := widget.NewEntryWithData(dnustr)

	workv, _ := zacatekstr.Get()
	workdnu, _ := dnustr.Get()

	vysledekstr.Set(doruceni(string_to_time(workv), string_to_int(workdnu)))

	vysledek := widget.NewLabelWithData(vysledekstr)

	w.SetContent(container.NewVBox(
		title,
		test,
		zacatek,
		dnu,
		vysledek,
		/*widget.NewButton("Spočítat", func() {
			vysledek.SetText(doruceni(string_to_time(zacatek.Text), string_to_int(dnu.Text)))
		}*/
	))

	w.ShowAndRun()
}
