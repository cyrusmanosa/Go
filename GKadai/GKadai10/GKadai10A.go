package main

import (
	"fmt"
	"os"
)

func main() {
	var age, person int
	fmt.Println("そうだ!動物園へ行こう!!")
	fmt.Print("何人で行きますか?> ")
	fmt.Scan(&person)
	fmt.Print("あなたの年齢を入力してください> ")
	fmt.Scan(&age)

	if age < 6 {
		fmt.Println("無料です!")
		os.Exit(0)
	} else if person < 30 {
		switch {
		case age > 6 && age < 15:
			fmt.Println("子供料金:200 円になります!")
		case age > 15:
			fmt.Println("大人料金:500 円になります!")
		}
	} else if person > 30 && person < 50 {
		switch {
		case age > 6 && age < 15:
			fmt.Println("子供料金:180 円になります!")
		case age > 15:
			fmt.Println("大人料金:450 円になります!")
		}
	} else if person > 50 {
		switch {
		case age > 6 && age < 15:
			fmt.Println("子供料金:160 円になります!")
		case age > 15:
			fmt.Println("大人料金:400 円になります!")
		}
	}
}
