package main

import "fmt"

func main() {
	var num int64
	fmt.Print("0 から 65535 までの整数を入力してください> ")
	fmt.Scan(&num)
	fmt.Println("最上位ビットから順に表示します!")
	fmt.Printf("%b", num)
}
