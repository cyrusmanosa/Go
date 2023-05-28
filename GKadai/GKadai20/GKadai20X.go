package main

import "math/rand"

func main() {
	println("バブルソートを使ってカードを昇順に並べます!")
	array := [10]int{}
	for j := 0; j < len(array); j++ {
		array[j] = rand.Intn(100)
	}

	for j := 0; j <= 45; j++ {
		print(j, "回目\t")
		for i := 0; i < len(array); i++ {
			if i+1 != len(array) {
				if array[i] > array[i+1] {
					c := array[i]
					array[i] = array[i+1]
					array[i+1] = c
				}
			}
			print(array[i], " ")
		}
		println(" ")
	}
}
