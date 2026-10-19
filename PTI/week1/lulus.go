package main

import "fmt"

func main() {
	var nilai int

	fmt.Scan(&nilai)

	if nilai >= 60 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Belum lulus")
	}

}
