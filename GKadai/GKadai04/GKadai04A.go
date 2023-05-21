package main

import "fmt"

func main() {
	var num int
	fmt.Print("教室番号を入力してください> ")
	fmt.Scan(&num)
	fmt.Println("一の位\t", num%10)
	fmt.Println("十の位\t", (num/10)%10)
	fmt.Println("百の位\t", (num/100)%10)
	fmt.Println("千の位\t", num/1000)
}
