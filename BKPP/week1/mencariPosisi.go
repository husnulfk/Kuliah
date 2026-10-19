package main

import "fmt"

func main() {
	var posisi0, s, t, posisi int

	fmt.Scan(&posisi0, &s, &t)

	posisi = posisi0 + (s * t)

	fmt.Println(posisi)
}
