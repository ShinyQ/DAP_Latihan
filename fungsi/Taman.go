/*
	Nama : Kurniadi Ahmad Wijaya
	NIM : 1301194024
*/

package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) float64 {
	var (
		a, b float64
	)

	a = math.Pow(math.Abs(x1-x2), 2)
	b = math.Pow(math.Abs(y1-y2), 2)
	return math.Sqrt(a + b)
}

func main() {
	var (
		x1, y1, x2, y2 float64
	)

	fmt.Print("(Titik1X, Titik1Y) : ")
	fmt.Scanln(&x1, &y1)

	fmt.Print("(Titik2X, Titik2Y) : ")
	fmt.Scanln(&x2, &y2)

	fmt.Printf("Jarak titik1 ke titik2 adalah : %.2f \n", jarak(x1, y1, x2, y2))

	fmt.Println("Nama : Kurniadi Ahmad Wijaya")
	fmt.Println("NIM : 1301194024")
}
