package main

import "fmt"

func main() {
	var (
		input int
	)

	fmt.Scanln(&input)

	for b := 1; b <= input; b++ {
		for c := input; c >= b; c-- { //spasi
			fmt.Print(" ")
		}
		for d := 1; d <= b; d++ { //bintang
			fmt.Print("*")
		}
		fmt.Println()
	}
}
