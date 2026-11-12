package main

import "fmt"

func main() {
	var belanja, total int
	var kartu, diskon, cashback bool

	fmt.Scan(&belanja, &kartu)

	if belanja >= 100000 && belanja < 200000 {
		diskon = true
		cashback = false
		total = belanja * 90 / 100
	} else if belanja >= 200000 && kartu {
		diskon = true
		cashback = true
		total = (belanja * 90 / 100) - 75000
	}
	fmt.Println("Kartu?", kartu, "Diskon?", diskon, "Cashback?", cashback, "Rp", total)
}
