package main

import (
	"fmt"
	"os"
)

func main() {
	var num, year int
	fmt.Println("こんにちは!")
	fmt.Println("わたしはスーパーティーチャー、ECC エクセレントよ!")
	fmt.Print("あなたのコースを教えてね!(1:IE、2:SK、3:SE)> ")
	fmt.Scan(&num)
	if num > 3 || num < 0 {
		fmt.Println("ちゃんと答えてね!")
		os.Exit(0)
	}

	fmt.Print("いま何年生? ")
	fmt.Scan(&year)
	switch num {
	case 1:
		if year < 0 || year > 4 {
			fmt.Println("IE は 4 年制コースよ!")
		}
	case 2:
		if year < 0 || year > 4 {
			fmt.Println("SK は 3 年制コースよ!")
		}
	case 3:
		if year < 0 || year > 4 {
			fmt.Println("SE は 2 年制コースよ!")
		}

	}

	if year > 0 || year < 4 {
		fmt.Println("IT カレッジの学生さん、よろしくね!")
	}
}
