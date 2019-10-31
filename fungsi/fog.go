/*
	Nama : Kurniadi Ahmad Wijaya
	NIM : 1301194024
	Deskripsi : Program fungsi komposisi (fogoh) (x) merupakan sebuah komposisi
	yang terdiri atas 4 rumus yaitu fungsi F dengan rumus x*x , lalu fungsi G
	dengan rumus x-2, kemudian fungsi H dengan rumus x+1. Dari ketiga rumus tadi
	digabungkan pada satu fungsi yang ada yaitu F(x) - (G(x) + H(x) sehingga
	dapat menghasilkan nilai fungsi komposisi (fogoh) (x)
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

	fmt.Println("")
	fmt.Println("Nama : Kurniadi Ahmad Wijaya")
	fmt.Println("NIM : 1301194024")
}
