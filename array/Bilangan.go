package main

import "fmt"

func main() {
	var (
		a           [10]int // [0 = 10, 1 = 20 , 2 , 3, 4, 5, 6, 7, 8, 9 = 30]
		N, i, input int
		max, min    int
	)
	i = 0
	N = len(a)  // N = 10
	for i < N { // 9 < 10; 1<10
		fmt.Print("Bilangan Ke ", i+1, " : ")
		fmt.Scanln(&input) // 10, 20
		a[i] = input       // a[0] = 10, a[1]= 20 ... a[9] = 30
		i++
	}
	i = 0

	fmt.Println("")
	fmt.Print("Bilangan : ")
	for i < N {
		fmt.Print(a[i]) // a[0] = 10; a[1] = 20; a[3] =
		fmt.Print(" ")
		i++
	}

	i = 0
	max = a[0]
	min = a[0]

	for i < N { // i = 1 , N = 10
		if a[i] < min { // a[1] = 20 < 10, min = 10
			min = a[i]
		}

		if a[i] > max { // a[1] = 20 > 10, max = 20
			max = a[i]
		}
		i++
	}
	fmt.Println("")
	fmt.Println("Max : ", max)
	fmt.Println("Min : ", min)

}
