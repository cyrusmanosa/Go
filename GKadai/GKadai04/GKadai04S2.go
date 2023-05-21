package main

import (
	"fmt"
)

func main() {
	var n, ans int
	fmt.Print("0 から 255 までの整数を入力してください> ")
	fmt.Scan(&n)

	fmt.Println("最下位ビットから順に表示します!")
	for i := 7; i >= 0; i-- {
		ans = n & 1
		if ans == 1 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
		n >>= 1
	}
}
