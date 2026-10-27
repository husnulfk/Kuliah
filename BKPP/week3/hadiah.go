package main

import (
	"fmt"
)

func main() {

	var n, total int
	var hadiah bool

	for fmt.Scan(&n); n != -1; fmt.Scan(&n) {
		total += n
	}
	hadiah = (total >= 10)

	fmt.Println(hadiah)
}
