package main

import "fmt"

func main() {
	var (
		w1, w2, w3, w4, input string
		isBerhasil            bool
	)

	isBerhasil = true

	for i := 1; i <= 5; i++ {

		fmt.Print("Percobaan ", i, ": ")
		fmt.Scanln(&w1, &w2, &w3, &w4)
		input = w1 + " " + w2 + " " + w3 + " " + w4
		isBerhasil = input == "merah kuning hijau ungu" && isBerhasil
	}

	fmt.Println("BERHASIL: ", isBerhasil)

}
