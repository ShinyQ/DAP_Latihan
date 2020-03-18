package main

import "fmt"

func main() {
	var (
		bil1, bil2, selisih int
	)

	fmt.Scanln(&bil1, &bil2)

	if bil1 < bil2 {
		selisih = (bil2 - bil1)
	} else {
		selisih = (bil1 - bil2)
	}

	fmt.Println(selisih)
}
