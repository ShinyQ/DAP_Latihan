package main

import "fmt"

func main() {
	var (
		input int
	)
	fmt.Scanln(&input)

	if input < 0 {
		fmt.Println("Negatif")
	} else if input > 0 {
		fmt.Println("Positif")
	} else {
		fmt.Println("Nol")
	}
}
