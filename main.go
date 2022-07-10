package main

import (
	"fmt"
)

var c = 34

func main() {
	x := "Hello World"
	fmt.Println(x)
	fmt.Println("Amaka favour is a wonderful lady. \nShe is a beautiful girl.")
	fmt.Println("1 + 1 =", 1+1)
	y := 1.0 + 1.0
	fmt.Println("1 + 1 =", y)
	z := 321325 * 424521
	fmt.Println("The product of 321325 * 424521 =", z)
	fmt.Println(c)
	me()
}

func me() {
	fmt.Println(c)

}
