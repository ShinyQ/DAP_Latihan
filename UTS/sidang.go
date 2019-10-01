package main

import "fmt"

func main() {
	var (
		input, MIN, i int
	)
	i = 1

	for input != -999 {

		fmt.Print("Mahasiswa ", i, " : ")
		fmt.Scanln(&input)

		if input >= 0 && input <= 700 {
			if i == 1 {
				MIN = input
			}

			if input < MIN {
				MIN = input
			}
			i++
		}
	}

	fmt.Println("SELESAI")
	fmt.Println("JUMLAH MAHASISWA : ", i-1)
	fmt.Println("MIN : ", MIN)
}
