package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	if x <= 1 {
		fmt.Println(false)
		return
	}

	isPrime := true

	for i := 2; i < x; i++ {

		if x%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)

}
