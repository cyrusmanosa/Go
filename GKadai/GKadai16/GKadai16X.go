package main

import "fmt"

func main() {
	heights := [7]int{160, 155, 170, 150, 175, 180, 165}
	println("一番背の高い人を一番後ろにします!")
	for i := 0; i < len(heights)-1; i++ {
		for j := i + 1; j < len(heights); j++ {
			if heights[i] > heights[j] {
				h := heights[j]
				heights[j] = heights[i]
				heights[i] = h
			}
		}
	}
	for key, value := range heights {
		fmt.Println("heights[", key, "]:", value)
	}
}
