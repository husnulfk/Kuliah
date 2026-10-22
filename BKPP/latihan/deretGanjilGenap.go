package main

import "fmt"

func main() {
	var n, ganjil, genap int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			genap++
		} else {
			ganjil++
		}
	}
	fmt.Printf("ganjil : %d, genap : %d\n", ganjil, genap)
}
