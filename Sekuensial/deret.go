package main

import "fmt"

func main() {
	var (
		phi, angka float64
		N, i       int
	)

	fmt.Print("N suku pertama :")
	fmt.Scanln(&N)

	i = 1
	angka = 1
	phi = 0

	for i <= N {
		if i%2 == 0 {
			phi = phi - (1.0 / angka)
		} else {
			phi = phi + (1.0 / angka)
		}
		angka += 2
		i++
	}
	fmt.Print("Hasil PI :", phi*4)
}
