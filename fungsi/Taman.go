/*
	Nama : Kurniadi Ahmad Wijaya
	NIM : 1301194024
	Deskripsi :
*/

package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) (total float64) {
	var (
		a, b float64
	)

	a = math.Sqrt(math.Abs(x1 - x2))
	b = math.Sqrt(math.Abs(y1 - y2))

	total = a + b
	return
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
