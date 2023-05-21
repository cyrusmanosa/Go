package main

import "fmt"

func main() {
	var num, total int
	fmt.Print("いくつまで足し合わせますか? ")
	fmt.Scan(&num)
	total = (num + 1) * num / 2
	fmt.Println("1 から", num, "まで足し合わせると・・・", total, " になりました!")
}
