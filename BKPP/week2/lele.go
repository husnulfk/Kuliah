package main

import "fmt"

func main() {
	var panjangLele float64
	var jumlahBulan int

	panjangLele = 3.0
	jumlahBulan = 0

	for i := 1; i <= 14; i++ {
		panjangLele = panjangLele + 0.3
		jumlahBulan = jumlahBulan + 1
	}

	fmt.Println(panjangLele, jumlahBulan)

}
