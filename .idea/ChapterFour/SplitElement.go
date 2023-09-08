package main

import "fmt"

func main() {
	inputSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	newSize := 3

	result := splitSlice(inputSlice, newSize)

	for _, size := range result {
		fmt.Println(size)
	}

}
func splitSlice(slice []string, newSize int) [][]string {
	var result [][]string

	for i := 0; i < len(slice); i += newSize {
		end := +newSize
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])

	}
	return result
}

//import (
//	"fmt"
//)
//
//func main() {
//	inputSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
//	chunkSize := 2
//
//	result := splitSlice(inputSlice, chunkSize)
//	fmt.Println(result)
//
//}
//
//func splitSlice(slice []string, chunkSize int) [][]string {
//	var result [][]string
//
//	for i:=0; i <chunkSize; i++ {
//		var end [] string
//		for j := 1; j < len(slice); j+= chunkSize {
//			end = append(end,slice[j])
//			//end := i + chunkSize
//			//if end > len(slice) {
//			//	end = len(slice)
//			//}
//			result = append(result, end)
//		}
//	}
//
//		return result
//}
