/*
	Nama : Kurniadi Ahmad Wijaya
	NIM : 1301194024
	Deskripsi : Untuk membuat sebuah program yang dapat menentukan range suatu nilai maka
	kita dapat menggunakan percabangan sesuai dengan range nilai pada deskripsi soal
*/

package main

import "fmt"

func main() {
	var (
		nam float64
		nmk string
	)
	fmt.Print("Nilai Akhir Mata Kuliah : ")
	fmt.Scanln(&nam)
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 && nam <= 80 {
		nmk = "AB"
	} else if nam > 65 && nam <= 72.5 {
		nmk = "B"
	} else if nam > 57.5 && nam <= 65 {
		nmk = "BC"
	} else if nam > 50 && nam <= 57.5 {
		nmk = "C"
	} else if nam > 40 && nam <= 50 {
		nmk = "D"
	} else {
		nmk = "E"
	}
	fmt.Println("Nilai Mata Kuliah :", nmk)
	fmt.Println("Nama : Kurniadi Ahmad Wijaya")
	fmt.Println("NIM : 1301194024")
}
