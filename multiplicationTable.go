package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var userInput int
	fmt.Scanf("%d", &userInput)
	multiplicationTable(userInput)

}
func multiplicationTable(number int) {
	for i := 0; i <= 12; i++ {
		fmt.Println(number, "*", i, "=", number*i)
	}
}
