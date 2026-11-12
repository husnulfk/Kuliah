package main

import "fmt"

func main() {
	var n, hasil int

	fmt.Scan(&n)

	if n <= 0 {
		hasil = n * -1
	} else {
		hasil = n
	}
	fmt.Println(hasil)
}
