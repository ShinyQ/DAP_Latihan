package main

import "fmt"

func main() {
	var (
		N, sjam, jam, menit, detik, smenit int
	)

	fmt.Scanln(&N)
	sjam = N / 3600
	jam = N % 3600
	menit = jam / 60
	smenit = jam % 60
	detik = smenit

	fmt.Println("Konversi ke jam, menit, dan detik :", sjam, "Jam", menit, "Menit", detik, "Detik ")
}
