package main

import "fmt"

func main() {
	var n, jumlah int

	fmt.Scan(&n)
	jumlah = 1

	for i := 0; i < n; i++ {
		jumlah = jumlah * (n - i)
	}
	fmt.Println(jumlah)
}
