package main

import "fmt"

func main() {
	f, r := 0, 0
	array := [5][5]string{
		{"ピタゴラス", "アルキメデス", "ユークリッド", "エラトステネス", "フィボナッチ"},
		{"デカルト", "フェルマー", "パスカル", "オイラー", "ラプラス"},
	}
	for {
		fmt.Print("階数を入力してください(0~1)> ")
		fmt.Scan(&f)
		fmt.Print("部屋番号を入力してください(0~4)> ")
		fmt.Scan(&r)
		if f > 1 || f < 0 || r > 4 || r < 0 {
			continue
		}
		fmt.Printf("%v階の%v号室には%vがいます\n", f, r, array[f][r])
		fmt.Println(" ")
	}
}
