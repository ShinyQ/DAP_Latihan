package main

import "fmt"

func main() {
	var (
		i, Batas, N, Nilai, Total int
		Rata                      float64
	)
	fmt.Print("Jumlah Calon Wisudawan: ")
	fmt.Scanln(&N)
	Batas = 1
	i = 1

	for i <= N {
		fmt.Print("Calon ", i, " EPrT ", Batas, " : ")
		fmt.Scanln(&Nilai)
		if Nilai >= 500 {
			Total += Nilai
			i++
			Batas = 1
		} else if Batas < 3 {
			Batas++
		} else {
			Total += Nilai
			Batas = 1
			i++
		}
	}

	Rata = float64(Total / N)
	fmt.Printf("Rerata EPrT : %.1f", Rata)
}
