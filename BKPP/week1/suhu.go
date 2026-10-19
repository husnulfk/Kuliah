package main

import "fmt"

func main() {
	var celsius, reamur, fahrenheit, kelvin float64
	fmt.Scan(&celsius)
	reamur = celsius * 4 / 5
	fahrenheit = celsius*9/5 + 32
	kelvin = celsius + 273.15

	fmt.Println(reamur, fahrenheit, kelvin)

}
