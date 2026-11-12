package main

import "fmt"

func main() {
	var tak float64

	fmt.Scan(&tak)

	if tak < 2.00 {
		fmt.Println("Poor")
	} else if tak >= 2.00 && tak <= 2.75 {
		fmt.Println("Fair")
	} else if tak >= 2.76 && tak <= 3.00 {
		fmt.Println("Satisfactory")
	} else if tak >= 3.01 && tak < 3.50 {
		fmt.Println("Very Good")
	} else {
		fmt.Println("Execelent")
	}
}
