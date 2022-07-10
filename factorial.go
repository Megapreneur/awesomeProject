package main

import "fmt"

func main() {
	factorial(5)

}
func factorial(number int) {
	result := 1
	for i := 1; i <= number; i++ {
		result *= i

	}
	fmt.Println("The factorial of", number, "is ", result)
}
