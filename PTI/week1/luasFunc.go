package main

import "fmt"

func main() {
	fmt.Println("Luas Persegi :", hitungLuas(3, 5))
}

func hitungLuas(panjang, lebar int) int {
	return panjang * lebar
}
