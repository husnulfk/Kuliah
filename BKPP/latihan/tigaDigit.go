package main

import (
	"fmt"
)

func main() {

	var x, d1, d2, d3 int

	fmt.Scan(&x)
	d1 = x / 100
	d2 = (x % 100) / 10
	d3 = x % 10
	fmt.Printf("d1 : %d  d2 : %d d3 : %d ", d1, d2, d3)
}
