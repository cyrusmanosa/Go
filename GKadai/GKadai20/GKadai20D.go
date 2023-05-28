package main

import "fmt"

func main() {
	nagayaECC := [...]string{
		"ピタゴラス", "アルキメデス", "ユークリッド", "エラトステネス", "フィボナッチ"}

	for i := 0; i < len(nagayaECC); i++ {
		fmt.Println(i, "号室:", nagayaECC[i])
	}

}
