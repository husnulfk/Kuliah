package main

import "fmt"

func main() {
	var uang int
	fmt.Scan(&uang)
	fmt.Println(uang / 10000)
	uang = uang % 10000
	fmt.Println(uang / 5000)
	uang = uang % 5000
	fmt.Println(uang / 1000)
}
