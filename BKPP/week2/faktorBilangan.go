package main

import "fmt"

func main() {
	var n int
	var s bool

	fmt.Scan(&n)

	for d := 1; d <= n; d++ {
		s = (n%d == 0)
		fmt.Println(d, s)
	}
}
