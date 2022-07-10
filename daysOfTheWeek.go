package main

import "fmt"

func main() {
	daysOfWeek(7)

}
func daysOfWeek(day int) {
	if day > 0 && day < 8 {
		switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		}
	} else {
		fmt.Println("Unknown number")
	}

}
