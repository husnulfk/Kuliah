package main

import (
	"fmt"
)

func main() {
	var x, y int
	var hasil int

	fmt.Scan(&x, &y)
	hasil = (x * x) + (2 * y) + 3
	fmt.Println(hasil)
}
