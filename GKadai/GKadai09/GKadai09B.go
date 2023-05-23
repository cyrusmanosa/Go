package main

import "fmt"

func main() {
	var num1, num2, person int
	fmt.Print("ひとつめの整数を入力してください> ")
	fmt.Scan(&num1)
	fmt.Print("ふたつめの整数を入力してください> ")
	fmt.Scan(&num2)
	fmt.Print("誰が計算しますか?(1:のび太、1 以外:出木杉)> ")
	fmt.Scan(&person)
	if person == 1 {
		fmt.Println("のび太が計算します!")
		fmt.Println("「", num1, "÷", num2, "は", num1/num2, "です!」")
	} else {
		fmt.Println("出木杉が計算します!")
		fmt.Println("「", num1, "÷", num2, "は", float64(num1)/float64(num2), "です!」")
	}
}
