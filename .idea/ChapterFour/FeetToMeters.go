package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter the lenght of Feet: ")
	var keyboard float64
	fmt.Scanf("%f", &keyboard)

	meters := keyboard * 0.3048
	fmt.Printf("Feet:%.2f", meters)
}
