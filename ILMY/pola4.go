package main

import "fmt"

func main() {
	var (
		input, bil int
	)

	fmt.Scanln(&input)
	bil = 9
	for b := 1; b <= input; b++ {
		for c := input; c >= b; c-- { //spasi
			fmt.Print(" ")
		}
		for d := 1; d <= b; d++ { //bintang

			fmt.Print(bil)

			if bil == -1 {
				bil = 9
			}
			bil--
		}
		fmt.Println()

	}
}
