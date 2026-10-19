package main

import "fmt"

func main() {
	var panjang, lebar, luas int

	fmt.Scan(&panjang, &lebar)

	luas = panjang * lebar

	fmt.Println(luas)
}
