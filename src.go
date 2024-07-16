package main

import (
	"fmt"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/cz"
)

/*func main() {
	date := input_time("Enter date in format 12.4.2024: \n")
	delta := input_int("Enter number of days: \n")
	print(doruceni(date, delta))
}*/

func doruceni(date time.Time, delta int) string {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(cz.Holidays...)
	konec := date.AddDate(0, 0, delta)
	var before_konec time.Time = konec
	posunuto := false
	for !c.IsWorkday(konec) {
		posunuto = true
		konec = konec.AddDate(0, 0, 1)
	}
	if !posunuto {
		return f("Konec lhuty bude %d.%d.%d (%s)", konec.Day(), konec.Month(), konec.Year(), convert_weekday(konec.Weekday()))
	}
	weekday_before_konec := convert_weekday(before_konec.Weekday())
	is_holiday, _, name_of_holiday := c.IsHoliday(before_konec)
	if is_holiday {
		weekday_before_konec += f(", %s", name_of_holiday.Name)
	}
	return f("Konec lhuty měl být %d.%d.%d (%s), ale kvuli svátku nebo víkendu bude až %d.%d.%d (%s)", before_konec.Day(), before_konec.Month(), before_konec.Year(), weekday_before_konec, konec.Day(), konec.Month(), konec.Year(), convert_weekday(konec.Weekday()))
}

/*func input_time(prompt string) time.Time {
	var d, m, y int
	fmt.Print(prompt)
	fmt.Scanf("%d.%d.%d", &d, &m, &y)
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}*/

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

/*func input_int(prompt string) int {
	var i int
	fmt.Print(prompt)
	for {
		_, err := fmt.Scanf("\n%d", &i)
		if err != nil {
			fmt.Println("Invalid input: ", err)
			continue
		}
		break
	}
	return i
}

func print(input ...any) {
	fmt.Println(input...)
}
*/
/* func printf(inputstr string, input ...any) {
	fmt.Printf(inputstr+"\n", input...)
} */

func f(inputstr string, input ...any) string {
	return fmt.Sprintf(inputstr, input...)
}

/* func sortarr(input ...any) {
	if f("%T", input[0]) == "[]int" {
		sort.Ints(input[0].([]int))
	} else if f("%T", input[0]) == "[]float64" {
		sort.Float64s(input[0].([]float64))
	} else if f("%T", input[0]) == "[]string" {
		sort.Strings(input[0].([]string))
	} else {
		fmt.Println("Invalid input type")
	}
} */
