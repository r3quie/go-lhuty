package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*
	func update_pravnimoc(value bool, pravmoc *widget.Label, zacatek *widget.Entry, dnu *widget.Entry) {
		if value {
			var konec string
			var kil string
			var kil2 string
			fmt.Sscanf(doruceni(string_to_time(zacatek.Text), string_to_int(dnu.Text)), "%v\n%v (%v)", &kil, &konec, &kil2)
			pravmoc.SetText("Nabude právní moci dne\n" + strformat(string_to_time(konec).AddDate(0, 0, 1)))
		}
		if !value {
			pravmoc.SetText("")
		}
	}
*/
func main() {
	a := app.New()
	w := a.NewWindow("Kalkulačka lhůt")
	dnes := time.Now()

	title := widget.NewLabel("Kalkulačka lhůt")
	labeldat := widget.NewLabel("Zadejte datum")
	labeldelta := widget.NewLabel("Zadejte počet dnů")

	zacatek := widget.NewEntry()
	zacatek.PlaceHolder = timeInFormat(dnes)
	dnu := widget.NewEntry()
	dnu.PlaceHolder = "8"

	vysledek := widget.NewLabel(doruceni(dnes, 8, false))

	pravnimoc := widget.NewCheck("Vypočítat právní moc", func(value bool) {
		vysledek.SetText(doruceni(string_to_time(zacatek.Text), string_to_int(dnu.Text), value))
	})

	w.SetContent(container.NewVBox(
		title,
		labeldat,
		zacatek,
		labeldelta,
		dnu,
		vysledek,
		pravnimoc,
	))

	w.Resize(fyne.NewSize(600, 200))

	dnu.OnChanged = func(s string) {
		vysledek.SetText(doruceni(string_to_time(zacatek.Text), string_to_int(s), pravnimoc.Checked))
	}
	zacatek.OnChanged = func(ss string) {
		vysledek.SetText(doruceni(string_to_time(ss), string_to_int(dnu.Text), pravnimoc.Checked))
	}

	w.ShowAndRun()
}
