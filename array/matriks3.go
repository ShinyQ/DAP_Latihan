package main

import "fmt"

func main() {
	var (
		a [3][3]int
		N int
	)

	N = 1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = N
			N += 10
		}
	}

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			fmt.Print(a[i][j], " ")

		}

		fmt.Println("")
	}

}
