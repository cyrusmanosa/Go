package main

import "fmt"

func main() {
	var mask [5]int
	fmt.Print("のび太の点数を入力してください> ")
	fmt.Scan(&mask[0])
	fmt.Print("しずかちゃんの点数を入力してください> ")
	fmt.Scan(&mask[1])
	fmt.Print("ジャイアンの点数を入力してください> ")
	fmt.Scan(&mask[2])
	fmt.Print("スネ夫の点数を入力してください> ")
	fmt.Scan(&mask[3])
	fmt.Print("出木杉くんの点数を入力してください> ")
	fmt.Scan(&mask[4])

	for i := 1; i < len(mask); i++ {
		if mask[i-1] > mask[i] {
			ans := mask[i]
			mask[i] = mask[i-1]
			mask[i-1] = ans

		}
	}
	fmt.Println("二番目に高い点数は", mask[3], "点です!")
}
