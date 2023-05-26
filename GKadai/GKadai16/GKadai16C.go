package main

import "fmt"

func main() {
	num := 0
	fmt.Print("カードがいくつありますか？＞　")
	fmt.Scan(&num)
	fmt.Println("カードが", num, "枚あります!")
	fmt.Println("カードを順番に表示します!")
	var a [100]int
	for i := 0; i < num; i++ {
		a[i] = i + 1
		fmt.Println("cards[", i, "]:", a[i])
	}
}
