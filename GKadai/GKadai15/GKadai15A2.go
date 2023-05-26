package main

import "fmt"

func main() {
	var num int
	fmt.Print("1 より大きい整数を入力してください> ")
	fmt.Scan(&num)
	fmt.Printf("1 から%vまで加算します!\n", num)
	fmt.Printf("合計は%vです!!", integrate(num)+integrate(1))
}

func integrate(n int) int {
	if n == 1 {
		return 1
	}
	var ans int
	for i := 1; i <= n; i++ {
		ans += i
	}
	return ans - 1
}
