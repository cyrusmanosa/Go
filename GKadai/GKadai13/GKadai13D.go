package main

import (
	"fmt"
	"os"
)

func main() {
	var num int
	fmt.Print("0 以上の整数を入力してください> ")
	fmt.Scan(&num)
	for num > 0 {
		fmt.Print(num % 10)
		num /= 10
	}
	if num == 0 {
		fmt.Print("0")
	}
	os.Exit(0)
}
