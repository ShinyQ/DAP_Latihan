package main

import "fmt"

func main() {
	var (
		input int
	)

	fmt.Scanln(&input)

	for i := 1; i <= input; i++ {
		for j := input; j >= i; j-- { //spasi
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ { //bintang
			fmt.Print("*")
		}
		for l := 1; l <= i-1; l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
