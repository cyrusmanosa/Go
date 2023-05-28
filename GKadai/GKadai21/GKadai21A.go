package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := 0
	fmt.Print("生成する整数の桁数を入力してください(1~9)> ")
	fmt.Scan(&num)

	r, array := 9, [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	max := len(array) - 1

	for i := 0; i < num; i++ {
		random := rand.Intn(r)
		print(array[random])
		array[random] = array[max]
		r--
		max--
	}
}
