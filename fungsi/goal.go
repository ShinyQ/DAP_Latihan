package main

import "fmt"

func is_kelipatan5(angka int) bool {
	return angka%5 == 0
}

func is_kelipatan2(angka int) bool {
	return angka%2 == 0
}

func is_ganjil(angka int) bool {
	return angka%2 == 1 && angka%5 != 0
}

func main() {
	var (
		bilangan int
		a, b     int
	)

	a, b = 0, 0
	for i := 1; i <= 5; i++ {
		fmt.Print("Tendangan Ke-", i, " ")
		fmt.Scanln(&bilangan)

		if is_ganjil(bilangan) {
			a++
			fmt.Println("Tendangan tepat sasaran")
		} else if is_kelipatan2(bilangan) {
			b++
			fmt.Println("Tendangan terlalu ke kiri atau terlalu ke kanan")
		} else if is_kelipatan5(bilangan) {
			b++
			fmt.Println("Tendangan terlalu ke atas ")
		}
		fmt.Println("")
	}

	fmt.Println("Skor akhir :", a, " untuk Marcus ", b, " untuk De Gea")
	if a > b {
		fmt.Println("Superb, Marcus!")
	} else {
		fmt.Println("Well done, De Gea!")
	}
}
