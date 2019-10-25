/*
	Nama : Kurniadi Ahmad Wijaya
	NIM : 1301194024
	Deskripsi :
*/

package main

import "fmt"

func fungsiF(x float64) float64 {
	return x * x
}

func fungsiG(x float64) float64 {
	return x - 2
}

func fungsiH(x float64) float64 {
	return x + 1
}

func fungsiFoGoH(x float64) float64 {
	return fungsiF(x) - (fungsiG(x) + fungsiH(x))
}

func main() {
	var N float64

	fmt.Print("Masukkan nilai x = ")
	fmt.Scanln(&N)

	fmt.Printf("f(%.2f) = %.2f \n", N, fungsiF(N))
	fmt.Printf("f(%.2f) = %.2f \n", N, fungsiG(N))
	fmt.Printf("f(%.2f) = %.2f \n", N, fungsiH(N))
	fmt.Printf("(fogoh) (%.2f) = %.2f \n", N, fungsiFoGoH(N))

	fmt.Println("Nama : Kurniadi Ahmad Wijaya")
	fmt.Println("NIM : 1301194024")
}
