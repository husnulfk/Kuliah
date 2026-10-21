package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	for i := n; i <= m; i++ {
		fmt.Println(i)
	}
}
