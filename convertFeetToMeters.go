package main

import "fmt"

func main() {
	feetToMeters(15)
}
func feetToMeters(feet float64) {
	m := feet * 0.3048
	fmt.Println(feet, "feet is equal to ", m, "meters")

}
