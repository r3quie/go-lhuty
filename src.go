package main

import (
	"fmt"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/cz"
)

const dateformat string = "%d.%d.%d"

// returns formatted date in format "day.month.year"
func strformat(x time.Time) string {
	return f(dateformat, x.Day(), x.Month(), x.Year())
}

// returns string with date + delta days, if it's a holiday or weekend, it returns the next workday
func doruceni(date time.Time, delta int, value bool) string {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(cz.Holidays...)
	var bude string = ""

	if delta == 0 && c.IsWorkday(date) {
		if time.Since(date) > 0 {
			bude = "byl"
		}
		bude = "bude"
		if value {
			return f("Konec lhůty %s \n%s (%s)\nA právní moci na%s %s", bude, strformat(date), convert_weekday(date.Weekday()), bude, strformat(date.AddDate(0, 0, 1)))
		}
		return f("Konec lhůty %s \n%s (%s)", bude, strformat(date), convert_weekday(date.Weekday()))
	}

	konec := date.AddDate(0, 0, delta)

	var before_konec time.Time = konec

	posunuto := false
	svatek := false
	vikend := false

	svatek_or_vikend := ""

	for !c.IsWorkday(konec) {
		posunuto = true

		h, _, n := c.IsHoliday(konec)

		if h && svatek {
			svatek_or_vikend += f("a svátku '%s' ", n.Name)
		} else if h {
			svatek = true
			svatek_or_vikend += f("svátku '%s' ", n.Name)
		}
		if konec.Weekday() == 0 || konec.Weekday() == 6 {
			vikend = true
		}

		konec = konec.AddDate(0, 0, 1)
	}
	bude = "bude"
	if time.Since(konec) > 0 {
		bude = "byl"
	}
	if !posunuto {
		if value {
			return f("Konec lhůty %s \n%s (%s)\nA právní moci na%s %s", bude, strformat(konec), convert_weekday(konec.Weekday()), bude, strformat(konec.AddDate(0, 0, 1)))
		}
		return f("Konec lhůty %s \n%s (%s)", bude, strformat(konec), convert_weekday(konec.Weekday()))
	}

	weekday_before_konec := convert_weekday(before_konec.Weekday())
	is_holiday, _, name_of_holiday := c.IsHoliday(before_konec)

	if is_holiday {
		weekday_before_konec += f(", %s", name_of_holiday.Name)
	}
	if !svatek {
		svatek_or_vikend = "víkendu"
	}
	if svatek && vikend {
		svatek_or_vikend += "a víkendu"
	}
	if value {
		return f("Konec lhůty měl být %s (%s), ale kvůli %s %s až \n%s (%s)\nA právní moci na%s %s", strformat(before_konec), weekday_before_konec, svatek_or_vikend, bude, strformat(konec), convert_weekday(konec.Weekday()), bude, strformat(konec.AddDate(0, 0, 1)))
	}
	return f("Konec lhůty měl být %s (%s), ale kvůli %s %s až \n%s (%s)", strformat(before_konec), weekday_before_konec, svatek_or_vikend, bude, strformat(konec), convert_weekday(konec.Weekday()))
}

// converts string in format "day.month.year" to time.Time
func string_to_time(input string) time.Time {
	var d, m, y int
	fmt.Sscanf(input, "%d.%d.%d", &d, &m, &y)
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

// converts string to int using fmt.Sscanf
func string_to_int(input string) int {
	var i int
	fmt.Sscanf(input, "%d", &i)
	return i
}

// converts time.Weekday to string in Czech
func convert_weekday(input time.Weekday) string {
	switch input {
	case 0:
		return "Neděle"
	case 1:
		return "Pondělí"
	case 2:
		return "Úterý"
	case 3:
		return "Středa"
	case 4:
		return "Čtvrtek"
	case 5:
		return "Pátek"
	case 6:
		return "Sobota"
	}
	return ""
}

// returns formatted, literally just fmt.Sprintf
func f(inputstr string, input ...any) string {
	return fmt.Sprintf(inputstr, input...)
}
