package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("\tN\tN2\tN3\tN4")
	for index := 1; index <= 5; index++ {
		for newIndex := 1; newIndex < 5; newIndex++ {
			multiply := int(math.Pow(float64(index), float64(newIndex)))
			fmt.Printf("%9d", multiply)
		}
		fmt.Println()

	}

}
