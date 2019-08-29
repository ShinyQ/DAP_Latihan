package main

import "fmt"

func main() {
	var (
		input     int  // Inisialisasi Variable Untuk Inputan User
		isKabisat bool // Inisialisasi Variable Untuk Status Kabisat
	)

	fmt.Print("Tahun: ")
	fmt.Scanln(&input) // User Memasukkan Inputan Tahun

	if input%400 == 0 || input%4 == 0 { // Pengecekan inputan user apabila dia habis dibagi 400 atau habis dibagi 4
		isKabisat = true // apabila ia habis dibagi 400 atau habis dibagi 4 maka satatus kabisat adalah true
	} else {
		isKabisat = false // apabila ia tidak habis dibagi 400 atau tidak habis dibagi 4 maka satatus kabisat adalah false
	}

	fmt.Println("Kabisat:", isKabisat) // Mengeluarkan output status kabisat sesuai operasi habis dibagi
}
