package main

import (
	"fmt"
)

func main() {
	var (
		i, total, nilai int
	)

	i = 1
	total = 0

	for i <= 3 {
		fmt.Print("Input Nilai Kuis ", i, " : ")
		fmt.Scanln(&nilai)

		if nilai >= 75 {
			total += nilai
			i++
		}
	}
	fmt.Print("Rata-rata Kuis DAP : ", total/3)

}
