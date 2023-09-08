//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//
//	for  {
//		var totalMiles float64
//		var totalGallons float64
//		//var milesGallon float64
//		var counter float64
//		fmt.Print("Enter miles driven: ")
//		var miles float64
//		fmt.Scanf("%f", &miles)
//		fmt.Print("Enter gallons used, press -1 to end: ")
//		var gallon float64
//		fmt.Scanf("%f", &gallon)
//		totalMiles += miles
//		totalGallons += gallon
//		counter++
//
//		milesGallon := totalMiles/totalGallons
//
//		if miles == -1 && gallon == -1{
//			average := milesGallon / counter
//			fmt.Println(average)
//		}
//
//
//
//
//
//
//	}
//
//
//
//
//}

package main

import (
	"fmt"
)

func main() {
	var totalMiles float64
	var totalGallons float64
	var counter float64

	for {
		fmt.Print("Enter miles driven (or -1 to end): ")
		var miles float64
		_, index := fmt.Scanf("%f", &miles)

		if index != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		if miles == -1 {
			break
		}

		fmt.Print("Enter gallons used: ")
		var gallon float64
		_, index = fmt.Scanf("%f", &gallon)

		if index != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		if gallon == -1 {
			break
		}

		totalMiles += miles
		totalGallons += gallon
		counter++

		milesGallon := totalMiles / totalGallons
		fmt.Printf("Current average miles per gallon: %.2f\n", milesGallon)
	}

	if counter > 0 {
		milesGallon := totalMiles / totalGallons
		average := milesGallon / counter
		fmt.Printf("Final average miles per gallon: %.2f\n", average)
	} else {
		fmt.Println("No data entered.")
	}
}
