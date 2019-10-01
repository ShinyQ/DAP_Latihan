/*
	Nama : Kurniadi Ahmad Wijaya
	NIM : 1301194024

	Untuk mencari tahun kabisat maka ketentuannya adalah bilangan tahun kabisat
	tersebut harus habis dibagi 400 atau habis dibagi 4 dan tidak habis
	dibagi 100 maka operasi untuk memnentukan bilangan tersebut merupkan
	kabisat atau bukan menggunakan operasi habis dibagi sehingga dapat menentukan
	apakah bilangan tersebut kabisat atau bukan
*/

package main

import "fmt"

func main() {
	var (
		tahun     int
		isKabisat bool
	)

	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	isKabisat = tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0

	fmt.Println("Kabisat:", isKabisat)

	fmt.Println("Nama: Kurniadi Ahmad Wijaya")
	fmt.Println("NIM: 1301194024")
}
