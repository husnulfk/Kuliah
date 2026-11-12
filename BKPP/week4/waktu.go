package main

import "fmt"

func main() {
	var t, jt, mt, dt, sisa int

	fmt.Scan(&t)

	jt = t / 3600
	sisa = t % 3600
	mt = sisa % 3600 / 60
	dt = sisa % 60 / 1

	fmt.Println(jt, mt, dt)

}
