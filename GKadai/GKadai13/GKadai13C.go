package main

import (
	"fmt"
	"os"
)

func main() {
	var num, ans int
	fmt.Print("1 より大きい整数を入力してください> ")
	fmt.Scan(&num)
	if num <= 1 {
		os.Exit(0)
	}
	fmt.Println("1 から 10 まで加算します!")
	for i := 1; i <= num; i++ {
		ans += i
	}
	fmt.Println("合計は", ans, "です!!")
}
