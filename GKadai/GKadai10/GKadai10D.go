package main

import "fmt"

func main() {
	var age int
	fmt.Println("そうだ!動物園へ行こう!!")
	fmt.Print("あなたの年齢を入力してください> ")
	fmt.Scan(&age)

	switch {
	case age < 6:
		fmt.Println("無料です!")
	case age > 6 && age < 15:
		fmt.Println("子供料金:200 円になります!")
	case age > 15:
		fmt.Println("大人料金:500 円になります!")
	}
}
