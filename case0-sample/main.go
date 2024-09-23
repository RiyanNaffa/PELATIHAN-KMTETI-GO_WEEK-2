// CASE 0: Menghitung Panjang Sisi Miring Segitiga

package main

import (
	"fmt"
	"math"
)

func CountLength(a, b float64) float64 {
	c := math.Sqrt((a * a) + (b * b))
	return c
}

func main() {
	a := 2.0
	b := 3.0

	c := CountLength(float64(a), float64(b))
	fmt.Println("length of c:", c)
}
