package main

import "fmt"

func main() {
	var x, hasil float64
	fmt.Scan(&x)
	hasil = ((x * x * x) + 3*x) / ((x * x * x * x) - 3*(x*x) + 4)
	fmt.Println(hasil)
}
