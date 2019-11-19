package main

import "fmt"

const SIZE = 3

var (
	arrDataCumlaude    [SIZE]Data
	arrDataNonCumlaude [SIZE]Data
	arrData            [SIZE]Data
)

type Data struct {
	nama, nim string
	semester  int
	ipk       float64
}

func main() {
	isiData()
	hitungIPK()
	tampilData()
}

func isiData() {
	var nama, nim string
	var semester int
	var ipk float64

	for i := 0; i < SIZE; i++ {
		fmt.Scanln(&nama, &nim, &ipk, &semester)
		arrData[i].nama = nama
		arrData[i].nim = nim
		arrData[i].ipk = ipk
		arrData[i].semester = semester
	}
}

func hitungCumlaude(ipk float64, semester int) bool {
	var cumlaude bool

	if ipk >= 3.50 && semester <= 8 {
		cumlaude = true
	}

	return cumlaude
}

func hitungIPK() {
	j := 0
	k := 0

	for i := 0; i < SIZE; i++ {
		if hitungCumlaude(arrData[i].ipk, arrData[i].semester) {
			arrDataCumlaude[j] = arrData[i]
			j++
		} else {
			arrDataNonCumlaude[k] = arrData[i]
			k++
		}
	}
}

func tampilData() {
	for i := 0; i < SIZE; i++ {
		fmt.Println("")
		fmt.Println("Nama :", arrDataCumlaude[i].nama)
		fmt.Println("IPK :", arrDataCumlaude[i].ipk)
		fmt.Println("NIM :", arrDataCumlaude[i].nim)
		fmt.Println("Jumlah Semester :", arrDataCumlaude[i].semester)
		fmt.Println("")
	}
}
