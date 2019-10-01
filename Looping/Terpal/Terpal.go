/*
	NIM : 1301194024
	Nama : Kurniadi Ahmad Wijaya

	Untuk membuat user dapat menginputkan data sampai salah satu kantong terpal berisi 9 Kg atau lebih
	dibutuhkan perulangan dengan kondisi dimana salah satu kantong tidak boleh > 9 atau berat
	kantong < 9 sehingga perulangan akan bekerja hingga pengguna menginputkan berat 9 di salah
	satu kantong
*/

package main

import "fmt"

func main() {
	var (
		kantong_1, kantong_2 float64
	)

	kantong_1 = 1
	kantong_2 = 1

	for kantong_1 < 9 && kantong_2 < 9 {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&kantong_1, &kantong_2)
	}
	fmt.Println("Proses selesai.")
	fmt.Println("Nama : Kurniadi Ahmad Wijaya")
	fmt.Println("NIM : 1301194024")
}
