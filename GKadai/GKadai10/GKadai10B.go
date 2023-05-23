package main

import (
	"fmt"
	"os"
)

func main() {
	var age int
	fmt.Println("そうだ!水族館へ行こう!!")
	fmt.Print("あなたの年齢を入力してください> ")
	fmt.Scan(&age)

	if age < 0 {
		fmt.Println("次にちゃんと答えてね!")
		os.Exit(0)
	}

	switch {
	case age < 3:
		fmt.Println("無料です!")
	case age > 3 && age < 6:
		fmt.Println("幼児料金:600 円になります!")
	case age > 6 && age < 15:
		fmt.Println("子供料金:1200 円になります!")
	case age > 15 && age < 65:
		fmt.Println("大人料金:2400 円になります!")
	default:
		fmt.Println("シニア料金:2200 円になります!")
	}
}
