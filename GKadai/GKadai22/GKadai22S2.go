package main

import "fmt"

func main() {
	array := [2][3][5]string{
		{
			{"ピタゴラス", "アルキメデス", "ユークリッド", "エラトステネス", "フィボナッチ"},
			{"デカルト", "フェルマー", "パスカル", "オイラー", "ラプラス"},
			{"フーリエ", "ガウス"},
		},
		{
			{"ド・モルガン", "ブール", "リーマン", "ポアンカレ"},
			{"ラッセル", "ニュートン", "テイラー"},
		},
	}
	fmt.Println("メゾン ECC の住人たちです!")
	for k := 0; k < len(array); k++ {
		fmt.Println(k, "号館")
		for i := 0; i < len(array[k]); i++ {
			if array[k][i][0] == "" {
				break
			}
			fmt.Print(i, "階の住人:")
			for j := 0; j < len(array[k][i]); j++ {
				if array[k][i][j] == "" {
					break
				}
				fmt.Print(j, ":", array[k][i][j], " ")
			}
			fmt.Println(" ")
		}
	}
}
