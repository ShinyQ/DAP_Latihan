package main

import "fmt"

func main() {
	var (
		nilaiA, nilaiB, nilaiAz, nilaiBz float64
		rataA, rataB                     float64
		P                                int
	)

	nilaiA = 0
	nilaiB = 0

	fmt.Scan(&P)
	var i = 0
	var j = 0

	for j < P*3 {
		fmt.Scan(&nilaiA)
		nilaiAz = nilaiAz + nilaiA
		j++
	}
	fmt.Println("Nilai Az", nilaiAz)
	fmt.Println("i", i)
	rataA = nilaiAz / float64(j)

	fmt.Scan(&P)

	for i < P*3 {
		fmt.Scan(&nilaiB)
		nilaiBz = nilaiBz + nilaiB
		i++
	}
	fmt.Println("i", i)
	fmt.Println("Nilai Bz", nilaiBz)
	rataB = nilaiBz / float64(i)

	fmt.Println("Rata-Rata Ahmad :", rataA, "Ahmad Menang ?", rataA > rataB)
	fmt.Println("Rata-Rata Badrun :", rataB, "Ahmad Menang ?", rataB > rataA)
}
