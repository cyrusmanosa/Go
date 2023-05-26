package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("数当てゲーム!")
	num, min, max := 0, 0, 99
	ans := rand.Intn(100)
	for i := 1; ; i++ {
		fmt.Print(i, "回目:", min, "から", max, "までの数値を入力してください> ")
		fmt.Scan(&num)
		if num == ans {
			fmt.Println("正解しました!")
			os.Exit(0)
		}
		if min > num || max < num {
			i--
			continue
		}

		if num > ans {
			max = num
			fmt.Println(num, "より小さいです!")
		} else {
			min = num
			fmt.Println(num, "より大きいです!")
		}
	}
}
