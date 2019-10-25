package main

import "fmt"

func main() {

	var (
		R, G, B   float64
		Cb, Cr, Y float64
	)

	fmt.Print("Nilai RGB : ")
	fmt.Scanln(&R, &G, &B)

	Y = 16 + (65.738 * R / 256) + (129.057 * G / 256) + (25.064 * B / 256)

	Cb = 128 - (37.945 * R / 256) - (74.494 * G / 256) + (112.439 * B / 256)[  ]

	Cr = 128 + (112.439 * R / 256) - (94.154 * G / 256) - (18.285 * B / 256)

	// fmt.Println("Y:", Y, ", Cb:", Cb, ", dan Cr:", Cr)
	fmt.Printf("Y: %.3f , Cb: %.3f dan Cr: %.3f ", Y, Cb, Cr)
}
