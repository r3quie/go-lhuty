package main

import (
	"fmt"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/cz"
)

func main() {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(cz.Holidays...)
	t := inputdate("Enter date in format 12.4.2024: \n")
	print(t)
}

func doruceni() []*cal.Holiday {
	holidays := cz.Holidays
	return holidays
}

func inputdate(prompt string) time.Time {
	var d, m, y int
	fmt.Print(prompt)
	fmt.Scanf("%d.%d.%d", &d, &m, &y)
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}

func print(input ...any) {
	fmt.Println(input...)
}

func printf(inputstr string, input ...any) {
	fmt.Printf(inputstr+"\n", input...)
}

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
