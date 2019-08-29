// Tentukan Proses Dari Program Berikut

package main

import "fmt"

func main() {

	var (
		satu, dua, tiga string // Inisialisasi Variable
		temp            string // Inisialisasi Variable
	)

	fmt.Print("Masukkan Input String : ") // Memasukkan Input Untuk Variable satu
	fmt.Scanln(&satu)                     // input(1)

	fmt.Print("Masukkan Input String : ") // Memasukkan Input Untuk Variable dua
	fmt.Scanln(&dua)                      // input(2)

	fmt.Print("Masukkan Input String : ") // Memasukkan Input Untuk Variable tiga
	fmt.Scanln(&tiga)                     // input(3)

	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga) // Output awal = 1 2 3

	temp = satu // membuat variable temp bernilai variable satu, temp = satu, temp = 1
	satu = dua  // membuat variable satu bernilai variable dua, satu = dua, satu = 2
	dua = tiga  // membuat variable dua bernilai variable tiga, dua = tiga, dua = 3
	tiga = temp // membuat variable tiga bernilai variable temp, tiga = temp, tiga = 1

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga) // Output akhir = 2 3 1
}
