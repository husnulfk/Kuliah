package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("kelipatan 3 dan 5")
	} else if n%3 == 0 {
		fmt.Println("kelipatan 3")
	} else if n%5 == 0 {
		fmt.Println("kelipatan 5")
	} else {
		fmt.Println(" ")
	}
}
