package main

import "fmt"

func main() {
	var (
		i                     int
		w1, w2, w3, w4, input string
		isBerhasil            bool
	)

	i = 1
	isBerhasil = true

	for i <= 5 {

		fmt.Print("Percobaan ", i, ": ")
		fmt.Scanln(&w1, &w2, &w3, &w4)
		input = w1 + " " + w2 + " " + w3 + " " + w4

		if input == "merah kuning hijau ungu" {

			if isBerhasil == false {
				isBerhasil = false
			} else {
				isBerhasil = true
			}

		} else {
			isBerhasil = false
		}

		i++
	}

	fmt.Println("BERHASIL: ", isBerhasil)

}
