package main

import "fmt"

func main() {
	var n, sisi, luas, keliling int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&sisi)
		luas = sisi * sisi
		keliling = sisi * 4
		fmt.Println(luas, keliling)
	}
}
