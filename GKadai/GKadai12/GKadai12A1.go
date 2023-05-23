package main

import "fmt"

func main() {
	var num1, num2, ans, l int
	fmt.Print("ひとつめの整数を入力してください> ")
	fmt.Scan(&num1)
	fmt.Print("ふたつめの整数を入力してください> ")
	fmt.Scan(&num2)

	if num1 > num2 {
		l = num1
	} else {
		l = num2
	}

	for i := 2; i < l; i++ {
		if (num1%i == 0) && (num2%i == 0) {
			ans = i
		}
	}
	fmt.Println("最大公約数は", ans, "です!")
}
