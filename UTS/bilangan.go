package main

import "fmt"

func main() {
	var (
		N, i, j, k, d int
	)

	fmt.Scanln(&N)
	i = 1

	for i <= N {
		j = i
		k = 1
		d = 0
		for k < j {
			if j%k == 0 {
				d = d + 1
			}
			k = k + 1
		}
		fmt.Println(j, d == 2)
		i = i + 1
	}
}
