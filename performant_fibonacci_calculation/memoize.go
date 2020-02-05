package main

import (
	"fmt"
	"os"
	"strconv"
)

func Fibonacci(n int, memo []int) int {
	if memo[n] != 0 {
		return memo[n]
	}
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	result := Fibonacci(n-1, memo) + Fibonacci(n-2, memo)
	memo[n] = result
	return result
}
func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Fibnaci of %d = %d \n", arg, Fibonacci(arg, make([]int, arg+1)))
}
