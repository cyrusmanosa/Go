package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("1 より大きい整数 1 を入力してください> ")
	fmt.Scan(&num1)
	fmt.Print("1 より大きい整数 2 を入力してください> ")
	fmt.Scan(&num2)
	fmt.Println(num1, "÷", num2, "の商と余りを計算します!")
	dekisugi(num1, num2)
}

func dekisugi(n1, n2 int) {
	fmt.Println("出木杉くん、よろしく!")
	fmt.Println("そんなの簡単さ!")
	a := n1 / n2
	b := n1 % n2
	fmt.Println("商は", a, "、余りは", b, "だよ。")
}
