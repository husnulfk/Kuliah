package main

import "fmt"

func main() {
	var untung1, untung2 float64

	fmt.Scan(&untung1, &untung2)

	if untung1 < untung2 {
		fmt.Println("peningkatan sebesar", untung2-untung1)
	} else if untung1 > untung2 {
		fmt.Println("penurunan sebesar", untung1-untung2)
	} else {
		fmt.Println("tetap")
	}
}
