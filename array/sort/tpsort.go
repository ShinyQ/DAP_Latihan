package main

const MAXSIZE = 20192020

type RecType struct {
	count int
	size  int
}

type ArrType [MAXSIZE]RecType

func iSort(tab ArrType, nsize int) {
	for i := 1; i < nsize; i++ {
		j := i
		for j > 0 {
			if tab[j-1].count > tab[j].count {
				tab[j-1], tab[j] = tab[j], tab[j-1]
			}
			j = j - 1
		}
	}
}

func mSort(tab ArrType, nsize int) {
	var minID int

	for i := 0; i < nsize; i++ {
		minID = i
		for j := i; j < nsize; j++ {
			if tab[j].size > tab[minID].size {
				minID = j
			}
		}
		tab[i].size, tab[minID].size = tab[minID].size, tab[i].size
	}
}

func isFound(tab ArrType, n, v int) bool {
	var iKanan = n
	var iKiri = 0

	for iKiri < iKanan {
		iMid := (iKanan + iKiri) / 2
		if tab[iMid].count < v {
			iKiri = iMid + 1
		} else if tab[iMid].count > v {
			iKanan = iMid - 1
		} else {
			return true
		}
	}

	return false
}
