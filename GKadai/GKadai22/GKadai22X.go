package main

import "fmt"

func main() {
	round, num := 0, 0
	fmt.Print("組み合わせを求める金額を入力してください> ")
	fmt.Scan(&num)
	money := num

	for i := 0; i <= money/100; i++ {
		moneyA := money
		moneyA -= 100 * i
		for j := 0; j <= moneyA/50; j++ {
			moneyB := moneyA
			moneyB -= 50 * j
			c := moneyB / 10
			fmt.Printf("100 円硬貨:%v枚、50 円硬貨:%v枚、10 円硬貨:%v枚\n", i, j, c)
			round++
		}
	}
	fmt.Println(money, "円になる組み合わせは全部で", round, "通りです!")
}
