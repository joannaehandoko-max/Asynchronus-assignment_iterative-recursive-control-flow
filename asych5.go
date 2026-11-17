package main

import "fmt"

func countEvenDigits(number int) int {
	if number == 0 {
		return 0
	}
	if (number%10)%2 == 0 {
		return 1 + countEvenDigits(number/10)
	}
	return countEvenDigits(number / 10)
}

func main() {
	fmt.Println(countEvenDigits(48271))
	fmt.Println(countEvenDigits(1359))
}
