package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("ひとつめの整数を入力してください> ")
	fmt.Scan(&num1)
	fmt.Print("ふたつめの整数を入力してください> ")
	fmt.Scan(&num2)

	ans2 := num1 * num2 / GCD(num1, num2)

	fmt.Print("最小公倍数は", ans2, "です!")
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
