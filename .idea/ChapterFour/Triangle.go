package main

import "fmt"

func main() {
	fmt.Println("Enter a number between 1 and 10: ")
	var input float64
	fmt.Scanf("%f", &input)

	for i := 0; i < (int(input)); i++ {
		for x := 0; x < i; x++ {
			fmt.Print("*" + " ")
		}
		fmt.Println()

	}
}
