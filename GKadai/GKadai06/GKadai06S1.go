package main

import "fmt"

var ans1, ans2 int

func main() {
	fmt.Print("ひとつめの整数を入力してください> ")
	fmt.Scan(&ans1)
	fmt.Println(" ")
	fmt.Print("ふたつめの整数を入力してください> ")
	fmt.Scan(&ans2)

	fans1 := float64(ans1)
	fans2 := float64(ans2)

	Ta := ans1 / ans2
	Tb := fans1 / fans2
	fmt.Println("\nのび太が計算します!")
	fmt.Println("「", ans1, "÷", ans2, " は", Ta, "です!」\n")
	fmt.Println("出木杉くんが計算します!")
	fmt.Println("「", ans1, "÷", ans2, " は", Tb, "です!」")
}
