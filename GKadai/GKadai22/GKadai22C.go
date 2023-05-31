package main

import "fmt"

func main() {
	array := [3][5]string{
		{"ピタゴラス", "アルキメデス", "ユークリッド", "エラトステネス", "フィボナッチ"},
		{"デカルト", "フェルマー", "パスカル", "オイラー", "ラプラス"},
		{"フーリエ", "ガウス"},
	}
	fmt.Println("メゾン ECC の住人たちです!")
	for i := 0; i < len(array); i++ {
		fmt.Print(i, "階の住人:")
		for j := 0; j < len(array[i]); j++ {
			if array[i][j] == "" {
				break
			}
			fmt.Print(j, ":", array[i][j], " ")
		}
		fmt.Println(" ")
	}
}
