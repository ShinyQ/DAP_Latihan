package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		i, N, A1, A2, MAX, MIN int
		msg                    string
	)
	i = 1
	fmt.Scanln(&N)
	for i <= N {
		fmt.Scanln(&A1, &A2)
		if A1 > A2 {
			MAX = A1
			MIN = A2
		} else {
			MAX = A2
			MIN = A1
		}

		msg += "\n" + "Max : " + strconv.Itoa(MAX) + " Min : " + strconv.Itoa(MIN)
		i++
	}

	fmt.Println(msg)
}
