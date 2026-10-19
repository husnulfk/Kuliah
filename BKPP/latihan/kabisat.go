package main

import (
	"fmt"
)

func main() {
	var x int
	var hasil bool

	fmt.Scan(&x)
	hasil = (x%400 == 0) || (x%4 == 0) && (x%100 != 0)
	fmt.Println(hasil)
}
