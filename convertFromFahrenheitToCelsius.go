package main

import "fmt"

func main() {
	fahrenheitToCelsius(70)

}
func fahrenheitToCelsius(F float64) {
	C := (F - 32) * 5 / 9

	fmt.Println(C)

}
