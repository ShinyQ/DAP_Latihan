package main

import "fmt"

func main() {
	var (
		N     int
		Pola  int = 9
		Space string
	)

	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		for j := 0; j+i < N-1; j++ {
			fmt.Print(" ")
		}
		Space = ""
		for j := 0; j <= i; j++ {
			Space = fmt.Sprintf("%d%s", Pola, Space)
			if Pola == 0 {
				Pola = 9
			} else {
				Pola--
			}
		}
		fmt.Println(Space)
	}
}
