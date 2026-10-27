package main

import "fmt"

func main() {
	var tabungan, minggu int
	 

	fmt.Scan(&tabungan, &minggu)

	for i := 1; i <= minggu; i++ {
		total = total + tabungan + (i-1)*2500
	}
	fmt.Println(total)
}
