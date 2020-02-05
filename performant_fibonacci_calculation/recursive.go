package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	result := fib(n-2) + fib(n-1)
	return result
}
func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Fibnaci of %d = %d \n", arg, fib(arg))
}
