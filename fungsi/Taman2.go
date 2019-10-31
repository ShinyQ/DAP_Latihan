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

	return a + b
}

func dalamLingkaran(x, r float64) bool {
	return x < math.Pow(math.Abs(r), 2)
}

func main() {
	var (
		x1, y1, x2, y2, r float64
	)

	fmt.Print("(PusatX, PusatY), R : ")
	fmt.Scanln(&x1, &y1, &r)

	fmt.Print("(TitikX, TitikY) : ")
	fmt.Scanln(&x2, &y2)

	if dalamLingkaran(jarak(x1, y1, x2, y2), r) {
		fmt.Printf("Titik (%.2f, %.2f) berada di DALAM lingkaran", x2, y2)
	} else {
		fmt.Printf("Titik (%.2f, %.2f) berada di LUAR lingkaran", x2, y2)
	}

	fmt.Println("")
	fmt.Println("Nama : Kurniadi Ahmad Wijaya")
	fmt.Println("NIM : 1301194024")
}
