package main

import "fmt"

func main() {
	var (
		input string
		i, j  int
	)

	fmt.Scanln(&input, &j)

	i = 1
	for i <= j {
		fmt.Println(input)
		i++
	}

}
