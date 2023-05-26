package main

import "fmt"

func main() {
	heights := [7]int{160, 155, 170, 150, 175, 180, 165}
	println("一番背の高い人を一番後ろにします!")
	for i := 1; i < len(heights); i++ {
		if heights[i-1] > heights[i] {
			h := heights[i-1]
			heights[i-1] = heights[i]
			heights[i] = h
		}
	}
	for key, value := range heights {
		fmt.Println("heights[", key, "]:", value)
	}
}
