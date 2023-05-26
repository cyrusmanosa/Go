package main

import "math/rand"

func main() {
	var array [9]int
	i := 1
	for i < 10 {
		r := rand.Intn(9)
		if array[r] != 0 {
			continue
		} else {
			array[r] = i
			i++
		}
	}
	for j := 0; j < len(array); j++ {
		println("card[", j, "]:", array[j])
	}
}
