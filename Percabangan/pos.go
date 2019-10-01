package main

import "fmt"

func main() {
	var (
		berat, b1, b2, sisa, biaya int
	)

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)
	sisa = 0

	if berat%1000 == 0 {
		b1 = berat / 1000
		fmt.Println("Detail berat: ", b1, " kg")
		biaya = b1 * 10000
		fmt.Println("Detail Biaya : Rp.", biaya)
	} else if berat > 10000 {
		b1 = berat / 1000
		b2 = berat % 1000
		fmt.Println("Detail berat: ", b1, " kg", " + ", b2, "gr")
		biaya = b1 * 10000
	} else {
		b1 = berat / 1000
		b2 = berat % 1000
		fmt.Println("Detail berat: ", b1, " kg", " + ", b2, "gr")
		biaya = b1 * 10000
		if b2 < 500 {
			sisa = b2 * 15
		} else {
			sisa = b2 * 5
		}
		fmt.Println("Detail Biaya : Rp.", biaya, " + Rp. ", sisa)
	}
	fmt.Println("Total Biaya : Rp.", biaya+sisa)
}
