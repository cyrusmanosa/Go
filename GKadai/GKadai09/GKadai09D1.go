package main

import "fmt"

func main() {
	var A, B int
	fmt.Print("ひとつめの整数を入力してください> ")
	fmt.Scan(&A)
	fmt.Println()

	fmt.Print("ふたつめの整数を入力してください> ")
	fmt.Scan(&B)
	fmt.Println()
	if A > B {
		fmt.Print("ふたつの数字で大きい方は", A, "です！")
	} else {
		fmt.Print("ふたつの数字で大きい方は", B, "です！")
	}
}
