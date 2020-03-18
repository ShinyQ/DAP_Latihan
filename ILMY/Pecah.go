package main

import "fmt"

func main() {
	var (
		input string
	)

	fmt.Scanf("%s", &input)
	for i := 0; i < len(input); i++ {
		fmt.Printf("Karakter Ke-%d %c\n", i+1, input[i])
	}
}
