package main

import "fmt"

func main() {
	var n int
	var kata string

	fmt.Scan(&n)

	fmt.Scan(&kata)
	for i := 0; i < n; i++ {
		fmt.Println(kata)
	}
}
