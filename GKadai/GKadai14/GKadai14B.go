package main

import (
	"fmt"
)

func main() {
	var n, p int
	fmt.Println("直角二等辺三角形の面積を求めます!")
	fmt.Print("斜辺の長さを入力してください>")
	fmt.Scan(&n)
	fmt.Print("誰に計算させますか?(1:出木杉、2:しずか)> ")
	fmt.Scan(&p)
	fmt.Println(" ")
	if p == 2 {
		ans := sizuka(n)
		fmt.Println("直角二等辺三角形の面積は", ans, "ね!")
	}

}

func sizuka(num int) float32 {
	fmt.Println("直角二等辺三角形が 4 つで正方形になるから")
	fmt.Println("正方形の面積を求めて 4 で割ればいいんだわ!")
	正方形 := float32(num) * float32(num)
	fmt.Println("正方形の面積は", 正方形, "だから")
	return 正方形 / 4
}
