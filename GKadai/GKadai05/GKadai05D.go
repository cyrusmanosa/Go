package main

import "math"

func main() {
	var num int = 16

	for i := 0; i < 10; i++ {
		if i == 0 {
			num *= 1
		} else {
			num *= 16
		}
		println("byte型：", byte(num), "\tint型：", int(num))
		if num > math.MaxInt64 {
			break
		}
	}
}
