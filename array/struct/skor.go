package main

import "fmt"

func main() {
	var (
		Hasil        []string
		KlubA, KlubB string
	)
	fmt.Print("Klub A : ")
	fmt.Scanln(&KlubA)

	fmt.Print("Klub B : ")
	fmt.Scanln(&KlubB)

	tambahSkor(&Hasil, KlubA, KlubB)
	tampilHasil(Hasil)

	fmt.Println("Pertandingan selesai")
}

func tambahSkor(Hasil *[]string, KlubA, KlubB string) {
	var skor1, skor2 int

	for i := 0; skor1 >= 0 && skor2 >= 0; i++ {
		fmt.Print("Pertandingan ", i+1, " : ")
		fmt.Scan(&skor1, &skor2)

		if skor1 > skor2 && skor2 >= 0 && skor1 >= 0 {
			fillWinner(Hasil, KlubA)
		} else if skor1 < skor2 && skor2 >= 0 && skor1 >= 0 {
			fillWinner(Hasil, KlubB)
		} else if skor1 == skor2 && skor2 >= 0 && skor1 >= 0 {
			fillWinner(Hasil, "Skor Seri")
		}
	}
}

func fillWinner(Hasil *[]string, name string) {
	*Hasil = append(*Hasil, name)
}

func tampilHasil(Hasil []string) {
	for i, val := range Hasil {
		fmt.Println("Hasil", i+1, " : ", val)
	}
}
