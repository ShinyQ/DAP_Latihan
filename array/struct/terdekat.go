package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

var (
	points    [10000]Point
	numpoints int
)

func jarak(p1, p2 Point) float64 {
	return math.Sqrt(((p1.x - p2.x) * (p1.x - p2.x)) + ((p1.y - p2.y) * (p1.y - p2.y)))
}

func bacaTitik() {
	var (
		point1, point2 float64
	)
	numpoints = 0
	fmt.Scanln(&point1, &point2)
	for point1 != 0 || point2 != 0 {
		points[numpoints] = Point{point1, point2}
		numpoints++
		fmt.Scanln(&point1, &point2)
	}
}

func ambilTitikTerdekat(p1, p2 *Point) {
	var min = jarak(points[0], points[1])

	for i := 0; i < numpoints; i++ {
		for j := i + 1; j < numpoints; j++ {
			if jarak(points[i], points[j]) < min {
				min = jarak(points[i], points[j])
				*p1 = points[i]
				*p2 = points[j]
			}
		}
	}
}

func main() {
	var (
		p1, p2 Point
	)

	bacaTitik()
	ambilTitikTerdekat(&p1, &p2)
	fmt.Printf("Titik terdekat adalah (%.1f,%.1f) dan (%.1f,%.1f) dengan jarak %.1f", p1.x, p1.y, p2.x, p2.y, jarak(p1, p2))
	fmt.Println("\nNama : Kurniadi Ahmad Wijaya")
	fmt.Println("NIM : 1301194024")
}
