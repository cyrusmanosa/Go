package main

import "fmt"

func main() {
	n, array := 0, [5]int{57300, 25000, 10000, 5000, 100}

	for {
		println(" ")
		for i := 0; i < len(array); i++ {
			println(i+1, "位:", array[i])
		}
		print("あなたのスコアを入力してください> ")
		fmt.Scan(&n)

		for i := 0; i < len(array); i++ {
			if n > array[i] {

				for j := len(array) - 1; j > i; j-- {
					array[j] = array[j-1]
				}

				array[i] = n
				println(i+1, "位になります!")
				break
			}
		}

	}
}
