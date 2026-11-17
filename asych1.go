package main

import (
	"fmt"
)

func main() {

	total := 0
	for i := 1; i <= 8; i++ {
		if i%2 == 0 {
			total = total + (i * 2)
		} else {
			total = total - i
		}
		fmt.Println(total)
	}

}
