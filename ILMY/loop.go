package main

import "fmt"

func main() {
	var (
		N, N1 int
	)

	fmt.Scanln(&N)

	if N%2 == 0 {
		N1 = N / 2
		for i := 1; i <= N1; i++ {
			fmt.Print("389")
			fmt.Print("353")
		}
	} else {

		N1 = N - 1
		N1 = N1 / 2
		for i := 1; i <= N1; i++ {
			fmt.Print("389")
			fmt.Print("353")
		}
		fmt.Print("389")
	}
	fmt.Println(N)
}
