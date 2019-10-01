package main

import (
	"fmt"
)

func main() {
	var (
		i             int
		bunga, output string
	)
	output = ""
	i = 0
	for bunga != "SELESAI" {
		i++
		output = output + bunga + " - "
		fmt.Print("Bunga ", i, ": ")
		fmt.Scanln(&bunga)
	}

	fmt.Println("Pita :", output)
	fmt.Println("Bunga :", i-1)
}
