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
		output = output + bunga + " - "
		i++
	}

	fmt.Println("Pita : ", output)
}
