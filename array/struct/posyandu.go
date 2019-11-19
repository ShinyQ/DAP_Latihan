package main

import "fmt"

var N int

func hitungMinMax(arrBerat []float64, bMin, bMax *float64) {
	min := arrBerat[0]
	max := arrBerat[0]
	for i := 0; i < N; i++ {
		if arrBerat[i] < min {
			min = arrBerat[i]
		}

		if arrBerat[i] > max {
			max = arrBerat[i]
		}
	}

	*bMin = min
	*bMax = max
}

func rerata(arrBerat []float64) float64 {
	total := 0.0
	for i := 0; i < N; i++ {
		total += arrBerat[i]
		fmt.Println(total)
	}
	total /= float64(N)
	return total
}

func main() {
	var (
		bMin, bMax float64
	)
	fmt.Print("Masukkan banyak data berat balita : ")
	fmt.Scanln(&N)

	berat := make([]float64, N)

	for i := 0; i < N; i++ {
		baby := 0.0
		fmt.Print("Masukkan Berat Balita Ke-", i+1, " : ")
		fmt.Scanln(&baby)

		berat[i] = baby
	}
	hitungMinMax(berat, &bMin, &bMax)

	fmt.Println("Berat Balita Minimum : ", bMin, " kg")
	fmt.Println("Berat Balita Maximum : ", bMax, " kg")
	fmt.Println("Rerata Berat Balita : ", rerata(berat), " kg")
}
