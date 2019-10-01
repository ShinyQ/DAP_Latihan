package main

import "fmt"

func main() {
	var (
		beratA, beratB, dimensi, panjang, lebar, tinggi, total float64
	)

	fmt.Print("Informasi Paket : ")
	fmt.Scanln(&beratA, &panjang, &lebar, &tinggi)

	beratB = panjang * lebar * tinggi / 6000
	dimensi = panjang + 2*(lebar+tinggi)

	if panjang >= 150 || lebar >= 150 || tinggi >= 150 && dimensi <= 400 {
		fmt.Println("Paket tidak dikirim: Terlalu panjang")
	} else {
		if dimensi <= 400 {
			if beratB < beratA {
				total = beratA
			} else {
				total = beratB
			}
			fmt.Println("Berat terhitung : ", total)
		} else {
			fmt.Println("Paket tidak dikirim: Terlalu besar")
		}
	}
}
