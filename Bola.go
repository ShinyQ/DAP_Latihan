package main

import (
	"fmt"
	"math" // Import Math Library untuk operasi perpangkatan
)

func main() {
	var (
		r, volume, luas float64                // Inisialisasi variabel r, volume, luas
		phi             float64 = 3.1415926535 // Inisialisasi variabel phi
	)

	fmt.Print("Jejari = ")
	fmt.Scanln(&r) // input(r) , jari-jari

	volume = 4.0 / 3.0 * phi * math.Pow(r, 3) // Operasi hitung volume
	luas = 4 * phi * math.Pow(r, 2)           // Operasi hitung luas

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f", int(r), volume, luas) // Output Operasi hitung dan luas
}
