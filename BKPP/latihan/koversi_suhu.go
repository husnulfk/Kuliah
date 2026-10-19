package main

import "fmt"

func main() {
	var celcius, reamur, fahreinheit, kelvin float64

	fmt.Scan(&celcius)

	reamur = celcius * 4 / 5
	fahreinheit = celcius*9/5 + 32
	kelvin = celcius + 273.15

	fmt.Println(reamur, fahreinheit, kelvin)
}
