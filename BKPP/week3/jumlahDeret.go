package main

import "fmt"

func main() {
	var n, m int
	var jumlah float64

	fmt.Scan(&n, &m)

	for i := n; i <= m; i++ {
		jumlah = jumlah + 4/float64(i)
		fmt.Println(jumlah)
	}
}
