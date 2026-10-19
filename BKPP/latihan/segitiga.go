package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var hasil bool

	fmt.Scan(&a, &b, &c)
	hasil = (a+b > c) && (b+c > a) && (a+c > b)
	fmt.Println(hasil)
}
