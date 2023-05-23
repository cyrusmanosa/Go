package main

import "fmt"

func main() {
	var num int
	fmt.Print("1 より大きい整数を入力してください> ")
	fmt.Scan(&num)
	fmt.Println("1 から", num, "まで加算します!")

	n, ans := 1, 0
	for n <= num {
		ans += n
		n++
	}
	fmt.Println("合計は", ans, "です!!")
}
