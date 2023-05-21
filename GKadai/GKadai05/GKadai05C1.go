package main

import (
	"math"
)

func main() {
	var r float64 = 5
	const p = 3.14
	area := math.Sqrt(r) * p

	println(float32(area))
}
