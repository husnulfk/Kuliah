package main

import "fmt"

func main() {
	total := hitungTotal(5000, 3)
	kembalian := hitungKembalian(20000, total)

	fmt.Println("Total belanja :", total)
	fmt.Println("Kembalian :", kembalian)
}

func hitungTotal(harga, jumlah int) int {
	return harga * jumlah
}

func hitungKembalian(uang, total int) int {
	return uang - total
}
