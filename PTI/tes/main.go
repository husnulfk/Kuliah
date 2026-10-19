package main

import "fmt"

func main() {
	var x, y int
	var count int = 1

	for count <= 5 {
		fmt.Scan(&x, &y)

		if x > y {
			fmt.Println(x)
		} else {
			fmt.Println(y)
		}
		count++
	}
}
