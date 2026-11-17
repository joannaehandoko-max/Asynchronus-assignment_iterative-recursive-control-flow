package main

import "fmt"

func main() {
	n := 15

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("COMPUTATION")
		} else if i%3 == 0 {
			fmt.Println("ALGORITHM")
		} else if i%5 == 0 {
			fmt.Println("LOGIC")
		} else {
			fmt.Println(i)
		}
	}
}
