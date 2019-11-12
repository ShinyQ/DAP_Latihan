package main

const N = 2019

type RecType struct {
	f1 string
	f2 int
	f3 float64
}
type ArrType [N]RecType

func main() {}

func rmax(data ArrType) float64 {
	nMax := data[0].f3

	for _, val := range data {
		if val.f3 > nMax {
			nMax = val.f3
		}
	}
	return nMax
}

func imin(data ArrType) int {
	var index int
	nMin := data[0].f2

	for i, val := range data {
		if val.f2 < nMin && val.f2 != 0 {
			nMin = val.f2
			index = i
		}
	}

	return index
}

func found(data ArrType, key string) bool {
	for _, val := range data {
		if val.f1 == key {
			return true
		}
	}
	return false
}

func pos(data ArrType, key string) int {

	var iKanan = N
	var iKiri = 0

	for iKiri < iKanan {
		iMid := (iKanan + iKiri) / 2
		if data[iMid].f1 < key {
			iKiri = iMid + 1
		} else if data[iMid].f1 > key {
			iKanan = iMid - 1
		} else {
			return iMid
		}
	}

	return -1
}
