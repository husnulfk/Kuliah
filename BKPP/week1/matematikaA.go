package main

import "fmt"

func main() {
	var x, y, hasil int
	fmt.Scan(&x, &y)
	hasil = (1/(3*(x*x) + 10)) + (10 * y) + 7
	fmt.Println(hasil)

}
