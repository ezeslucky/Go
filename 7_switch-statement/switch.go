package main

import "time"

func main() {

	// simple switch statement
	switch day := 3; day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	default:
		println("Another day")
	}

	// switch with multiple cases
	switch month := 5; month {
	case 1, 2, 3:
		println("First quarter")
	case 4, 5, 6:
		println("Second quarter")
	case 7, 8, 9:
		println("Third quarter")
	default:
		println("Fourth quarter")
	}

	// or
	switch time.Now().Month() {
	case time.January, time.February, time.March:
		println("First quarter of the year")
	case time.April, time.May, time.June:
		println("Second quarter of the year")
	case time.July, time.August, time.September:
		println("Third quarter of the year")
	default:
		println("Fourth quarter of the year")
	}

	//type switch
	var i interface{} = "Hello"
	switch v := i.(type) {
	case string:
		println("String:", v)
	case int:
		println("Int:", v)
	default:
		println("Unknown type")
	}
}