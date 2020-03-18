package main

import "fmt"

func main() {
	var (
		a, b, n int
	)

	fmt.Scanln(&n, &a, &b)

	for i := 0; i < n; i++ {
		fmt.Println(a + (b * i))
	}
}
