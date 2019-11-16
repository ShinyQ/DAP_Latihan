package main

import "fmt"

var (
	TabelKlub []Klub
	m         int
)

type Klub struct {
	nama string
	poin int
}

func main() {

	fmt.Print("Masukkan Jumlah Klub : ")
	fmt.Scanln(&m)
	isiData(m)
	fmt.Println("")
	sortAsc()
	fmt.Println("")
	cariKlub()
	fmt.Println("")
	sortDesc()
}

func isiData(n int) {
	var (
		nama string
		poin int
	)
	for i := 1; i <= n; i++ {
		fmt.Print("Nama Klub ", i, ": ")
		fmt.Scanln(&nama)

		fmt.Print("Poin Klub ", i, ": ")
		fmt.Scanln(&poin)

		klub := Klub{
			nama: nama,
			poin: poin,
		}

		TabelKlub = append(TabelKlub, klub)
	}
}

func cariKlub() {
	var (
		nama    string
		Selesai bool
		data    int
	)
	for nama != "cukup" {
		fmt.Print("Cari klub: ")
		fmt.Scanln(&nama)

		for i := 0; i < len(TabelKlub) && Selesai != true; i++ {
			if TabelKlub[i].nama == nama {
				data = i
				Selesai = true
			} else {
				data = -1
			}
		}
		Selesai = false

		if nama != "cukup" {
			if data != -1 {
				fmt.Println("Klub", nama, "memperoleh", TabelKlub[data].poin, "poin")
			} else {
				fmt.Println("Klub", nama, "tidak ada dalam daftar")
			}
		}
	}

}

func sortAsc() {
	var (
		n      int
		sorted bool
	)

	dataKlub := TabelKlub
	n = len(dataKlub)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if dataKlub[i].nama > dataKlub[i+1].nama {
				dataKlub[i+1], dataKlub[i] = dataKlub[i], dataKlub[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n--
	}
	n = len(dataKlub)
	fmt.Println("Urutan Klub berdasarkan nama ascending")
	for i := 0; i < n; i++ {
		fmt.Println(i+1, ". Nama:", dataKlub[i].nama, "dengan poin:", dataKlub[i].poin)
	}
}

func sortDesc() {
	var (
		n      int
		sorted bool
	)

	dataKlub := TabelKlub
	n = len(dataKlub)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if dataKlub[i].poin < dataKlub[i+1].poin {
				dataKlub[i+1], dataKlub[i] = dataKlub[i], dataKlub[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n--
	}
	n = len(dataKlub)
	fmt.Println("Urutan Klub berdasarkan nama ascending")
	for i := 0; i < n; i++ {
		fmt.Println(i+1, ". Nama:", dataKlub[i].nama, "dengan poin:", dataKlub[i].poin)
	}
}
