package main

import (
	"fmt"
	"os"
	"strconv"
)

func CountSets(arr []int, total int) int {
	return Recursive(arr, total, len(arr)-1)
}
func Recursive(arr []int, total int, i int) int {
	if total == 0 {
		return 1
	} else if total < 0 {
		return 0
	} else if i < 0 {
		return 0
	} else if total < arr[i] {
		return Recursive(arr, total, i-1)
	}
	return Recursive(arr, total-arr[i], i-1) + Recursive(arr, total, i-1)
}
func main() {
	total, _ := strconv.Atoi(os.Args[1])
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
	fmt.Printf("Number of sets in given array are: %d\n", CountSets(values, total))
}
