package main

import "fmt"

func main() {
	var (
		input, hasil int
	)

	fmt.Scanln(&input)

	hasil = 0
	for i := 0; i < input; i++ {
		hasil += (i + 1)
		fmt.Println(i + 1)
	}

	fmt.Println(hasil)

}
