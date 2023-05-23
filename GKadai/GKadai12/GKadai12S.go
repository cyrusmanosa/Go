package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	println("数当てゲーム!")
	ran := rand.Intn(100)
	ans := 0

	for i := 1; ; i++ {
		fmt.Print(i, " 回目:数値を入力してください> ")
		fmt.Scan(&ans)
		switch {
		case ans > ran:
			fmt.Println(ans, "より小さいです!")
		case ans < ran:
			fmt.Println(ans, "より大きいです!")
		default:
			fmt.Println("正解しました!")
			os.Exit(0)
		}
	}
}
