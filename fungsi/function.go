package main

import "fmt"

func avg(x float64, y float64) (float64, float64) {
	if x > 5 {
		x = 10
		y = 10
	} else {
		x = 1
		y = 1
	}

	return x, y
}

func main() {
	var (
		x, y            float64
		result, result2 float64
	)
	i := 1

	for i <= 3 {
		fmt.Scanln(&x, &y)
		result, result2 = avg(x, y)
		fmt.Println("Hasil Penjumlahan : ", result, result2)
		fmt.Println("")
		i++
	}

}
