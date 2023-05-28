package main

import (
	"math/rand"
)

func main() {
	array := [10]int{}
	println("最初の並びを表示します!")
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
		print(array[i], " ")
	}
	println("\n逆順にします!")
	for i := 0; i < len(array)/2; i++ {
		min, max, c := 0, len(array)-1, array[i]
		array[min+i] = array[max-i]
		array[max-i] = c

		for i := 0; i < len(array); i++ {
			print(array[i], " ")
		}
		println(" ")
	}
}
