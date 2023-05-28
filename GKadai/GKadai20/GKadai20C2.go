package main

import "fmt"

func main() {
	num, ans := 0, 0
	fmt.Print("1 より大きい整数を入力してください> ")
	fmt.Scan(&num)
	fmt.Printf("1 から %v まで加算します!", num)
	for i := 1; i <= num; i++ {
		ans += i
	}
	fmt.Printf("\n合計は%vです!!", ans)
}
