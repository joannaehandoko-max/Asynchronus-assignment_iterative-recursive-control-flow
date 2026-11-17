package main

import "fmt"

func main() {
	readings := []float64{190, 200, 215, 210, 185}
	n := len(readings)

	ideal := 0
	overheated := 0
	total := 0.0

	for i := 0; i < n; i++ {
		temp := readings[i]
		total = total + temp
		if temp >= 195 && temp <= 210 {
			ideal = ideal + 1
		} else if temp > 210 {
			overheated = overheated + 1
		}
	}

	fmt.Println("Ideal Minutes =", ideal)
	fmt.Println("Overheated Minutes =", overheated)
	fmt.Println("Average Temperature =", total/float64(n))
}
