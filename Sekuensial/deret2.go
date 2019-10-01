/*
	Nama : Kurniadi Ahmad Wijaya
	NIM : 1301194024
	Deskripsi : Untuk membuat sebuah program dengan selisih antar dua suku maka dibuat sebuah
	variable untuk menampung nilai asli variable lama kemudian sebuah variable baru untuk
	menampung hasil phi suku selanjutnya dengan menggunakan perulangan dengan ketentuan nilai selisih
	antara phi awal dan akhir maka dapat kita hasilkan suku ke-i
*/

package main

import "fmt"

func main() {
	var (
		phi, phi1, selisih, N float64 = 0, 0, 1, 1
		i                     int     = 1
	)

	for selisih > 0.00001 {

		phi1 = phi

		if i%2 == 0 {
			phi -= 1 / N
		} else {
			phi += 1 / N
		}

		selisih = phi1*4 - phi*4
		N += 2

		if selisih < 0 {
			selisih *= -1
		}
		i++
	}

	fmt.Println("Hasil PI : ", phi1*4)
	fmt.Println("Hasil PI : ", phi*4)

	fmt.Println("Nama : Kurniadi Ahmad Wijaya")
	fmt.Println("NIM : 1301194024")
}
