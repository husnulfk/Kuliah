package main

import "fmt"

func main() {
	var belanja, potongan float64
	var clo bool

	fmt.Scan(&belanja, &clo)

	if clo {
		potongan = belanja * 0.35
	} else {
		potongan = 0
	}
	fmt.Println(belanja - potongan)
}
