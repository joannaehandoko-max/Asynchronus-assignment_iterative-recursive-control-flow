package main

import "fmt"

func fastPower(x int, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		half := fastPower(x, n/2)
		return half * half
	}
	return x * fastPower(x, n-1)
}

func main() {
	fmt.Println(fastPower(2, 6))
}
