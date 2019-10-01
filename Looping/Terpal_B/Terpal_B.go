/*
	NIM : 1301194024
	Nama : Kurniadi Ahmad Wijaya

	Untuk membuat kondisi true oleng saat selisih kedua kantong lebih atau sama dengan 9kg
	maka dibuat sebuah kondisi apabila kantong pertama dikurangi kantong kedua lebih dari
	sama dengan 9 atau kantong kedua dikurangi kantong awal lebih dari sama dengan 9
	sehingga dapat diketahui apakah akan oleng atau tidak untuk pemberhentian proses
	looping jika berat total kantong adalah 150kg atau lebih maka tinggal melakukan
	operasi penjumlahan saja dengan cara kantong pertama ditambah kantong kedua
	kemudian di kondisikan apakah ia kurang dari 150 atau tidak
*/
package main

import "fmt"

func main() {
	var (
		kantong_1, kantong_2 float64
		isOleng              bool
	)

	fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
	fmt.Scanln(&kantong_1, &kantong_2)

	for kantong_1+kantong_2 < 150 && kantong_1 >= 0 && kantong_2 >= 0 {
		isOleng = kantong_1-kantong_2 >= 9 || kantong_2-kantong_1 >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng: ", isOleng)

		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&kantong_1, &kantong_2)
	}

	fmt.Println("Proses selesai.")
	fmt.Println("Nama : Kurniadi Ahmad Wijaya")
	fmt.Println("NIM : 1301194024")
}
