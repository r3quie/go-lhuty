package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/cz"
)

// returns string with date + delta days, if it's a holiday or weekend, it returns the next workday
func doruceni(date time.Time, delta int, value bool) string {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(cz.Holidays...)
	var bude string
	if delta < 0 {
		delta = -delta
	}
	if delta == 0 && c.IsWorkday(date) {
		if time.Since(date) > 0 {
			bude = "byl"
		} else if time.Since(date) < 0 {
			bude = "bude"
		} else if time.Since(date) == 0 {
			bude = "je dnes,"
		}

		if value {
			return f("Poslední den lhůty %s \n%s (%s)\nA právní moci na%s %s",
				bude, timeInFormat(date), convert_weekday(date.Weekday()), bude, timeInFormat(date.AddDate(0, 0, 1)))
		}
		return f("Poslední den lhůty %s \n%s (%s)",
			bude, timeInFormat(date), convert_weekday(date.Weekday()))
	}

	konec := date.AddDate(0, 0, delta)

	before_konec := konec

	posunuto := false
	svatek := false
	vikend := false

	var svatek_or_vikend string

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
			return f("Poslední den lhůty %s \n%s (%s)\nA právní moci na%s %s",
				bude, timeInFormat(konec), convert_weekday(konec.Weekday()), bude, timeInFormat(konec.AddDate(0, 0, 1)))
		}
		return f("Poslední den lhůty %s \n%s (%s)",
			bude, timeInFormat(konec), convert_weekday(konec.Weekday()))
	}

	weekday_before_konec := convert_weekday(before_konec.Weekday())
	is_before_holiday, _, name_of_holiday := c.IsHoliday(before_konec)

	if is_before_holiday {
		weekday_before_konec += f(", %s", name_of_holiday.Name)
	}
	if !svatek {
		svatek_or_vikend = "víkendu"
	}
	if svatek && vikend {
		svatek_or_vikend += "a víkendu"
	}
	if value {
		return f("Poslední den lhůty měl být %s (%s), ale kvůli %s %s až \n%s (%s)\nA právní moci na%s %s",
			timeInFormat(before_konec), weekday_before_konec, svatek_or_vikend, bude, timeInFormat(konec), convert_weekday(konec.Weekday()), bude, timeInFormat(konec.AddDate(0, 0, 1)))
	}
	return f("Poslední den lhůty měl být %s (%s), ale kvůli %s %s až \n%s (%s)",
		timeInFormat(before_konec), weekday_before_konec, svatek_or_vikend, bude, timeInFormat(konec), convert_weekday(konec.Weekday()))
}

// converts string in format "day.month.year" (2.1.2006) to time.Time
func string_to_time(input string) time.Time {
	date, err := time.Parse("2.1.2006", input)
	if err != nil {
		return time.Now()
	}
	return date
}

// converts string to int using Atoi, if fails, returns 0
func string_to_int(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return i
}

// converts time.Weekday to corresponding string in Czech
func convert_weekday(weekday time.Weekday) string {
	weekdays := [7]string{"Neděle", "Pondělí", "Úterý", "Středa", "Čtvrtek", "Pátek", "Sobota"}
	return weekdays[weekday]
}

// returns formatted date in format "day.month.year" (2.1.2006) from time.Time
func timeInFormat(x time.Time) string {
	return x.Format("2.1.2006")
}

// returns formatted, literally just fmt.Sprintf
func f(inputstr string, input ...any) string {
	return fmt.Sprintf(inputstr, input...)
}
