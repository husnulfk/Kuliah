package main

import (
	"fmt"
)

func main() {
	var alphabet string

	fmt.Scan(&alphabet)

	if alphabet >= "a" && alphabet <= "z" || alphabet >= "A" && alphabet <= "Z" {
		fmt.Println("Alpahbet")
	} else {
		fmt.Println("Bukan Alphabet")
	}
}
