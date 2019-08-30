package main

import (
	"fmt"
)

func main() {
	var (
		n, i          int
		bunga, output string
	)

	fmt.Scanln(&n)
	i = 1

	for i <= n {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scanln(&bunga)

		if bunga == "SELESAI" {
			break
		} else {
			output = output + bunga + " - "
			i++
		}
	}

	fmt.Println("Pita :", output)
	fmt.Println("Bunga :", i-1)
}
