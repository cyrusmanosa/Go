package main

import "fmt"

func main() {
	room, num := [5]string{"ピタゴラス", "アルキメデス", "ユークリッド", "エラトステネス", "フィボナッチ"}, 0
	for {
		fmt.Print("何号室を見ますか? ")
		fmt.Scan(&num)
		if num > len(room)-1 {
			fmt.Println("そんな部屋番号はありません!")
		} else {
			fmt.Println(room[num], "が住んでいます!")
		}
	}
}
