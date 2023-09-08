package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter your fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celsius := fahrenheit - 32*0.56
	fmt.Println(celsius)
}
