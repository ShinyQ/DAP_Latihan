package main

import "fmt"

func main() {
	var (
		sn, g, s, r, t, w float64
		u, a              int
	)
	fmt.Scanln(&a)
	g = 1.618
	s = 2.236
	r = 1
	t = 0.618
	u = 0
	w = 1

	for a >= 0 {
		for u <= a {
			r = g * r
			w = t * w
			u++
		}

		sn = (r - w) / s
		fmt.Println(sn)
		fmt.Scanln(&a)
	}
	fmt.Println("Selesai")
}
