package main

import (
	"fmt"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/cz"
)

func main() {
	date := input_time("Enter date in format 12.4.2024: \n")
	delta := input_int("Enter number of days: \n")
	print(doruceni(date, delta))
}

func doruceni(inputdate time.Time, inputdelta int) string {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(cz.Holidays...)
	konec := inputdate.AddDate(0, 0, inputdelta)
	var before_konec time.Time = konec
	posunuto := false
	for !c.IsWorkday(konec) {
		posunuto = true
		konec = konec.AddDate(0, 0, 1)
	}
	if !posunuto {
		return f("Konec lhuty bude %d.%d.%d", konec.Day(), konec.Month(), konec.Year())
	}
	return f("Konec lhuty měl být %d.%d.%d, ale kvuli svátku nebo víkendu bude až %d.%d.%d", before_konec.Day(), before_konec.Month(), before_konec.Year(), konec.Day(), konec.Month(), konec.Year())
}

func input_time(prompt string) time.Time {
	var d, m, y int
	fmt.Print(prompt)
	fmt.Scanf("%d.%d.%d", &d, &m, &y)
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

func input_int(prompt string) int {
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
