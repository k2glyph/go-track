package main

import (
	"fmt"
	"os"
	"strconv"
)

// Bottom Up approach
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	bottomUp := make([]int, n+1)
	bottomUp[1] = 1
	bottomUp[2] = 1
	for i := 3; i <= n; i++ {
		bottomUp[i] = bottomUp[i-1] + bottomUp[i-2]
	}
	return bottomUp[n]
}
func main() {
	arg, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Fibnaci of %d = %d \n", arg, Fibonacci(arg))
}
