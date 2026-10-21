package main

import "fmt"

func main() {
	var n, jam, totalJam int
	var rata float64

	fmt.Scan(&n)
	totalJam = 0

	for i := 0; i < n; i++ {
		fmt.Scan(&jam) 
		totalJam = totalJam + jam
	}
	
	rata = float64(totalJam) / float64(n)
	fmt.Println(rata)
}
