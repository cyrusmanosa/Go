package main

import "fmt"

func main() {
	燃料, 速度, 高度 := 15, 0, 300
	c, i := 0, 1
	fmt.Println("月面着陸ゲーム!")
	fmt.Println("ロケットを操って無事着陸してください。")
	for {
		if 高度 < 0 && (速度 > -10 && 速度 < 0) {
			fmt.Println("おめでとう!着陸成功!!")
		} else if 高度 < 0 && 速度 < -10 {
			fmt.Println("ズキューン!着陸失敗!!")
		}
		fmt.Println("--------------------------------")
		fmt.Printf("[燃料]:%v\t[速度]:%v\t[高度]:%v", 燃料, 速度, 高度)
		fmt.Println("--------------------------------")
		fmt.Print("逆噴射しますか?(1:する、それ以外:しない)> ")
		fmt.Scan(&c)
		if c == 1 {
			燃料--
			速度++
			高度++
		} else {
			高度 = 高度 - (4 * i)
			速度--
			i++
		}
	}
}
