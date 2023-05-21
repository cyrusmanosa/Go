package main

import (
	"fmt"
)

func main() {
	var n, ans int
	fmt.Print("0 から 255 までの整数を入力してください> ")
	fmt.Scan(&n)

	fmt.Println("最下位ビットから順に表示します!")
	for n > 0 {
		ans = n & 1
		if ans == 1 {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
		n >>= 1
	}
}
