package main

import "fmt"

func main() {
	fmt.Println("こんにちは!")
	fmt.Println("わたしはスーパーティーチャー、ECC エクセレントよ!")
	fmt.Println(" ")
	fmt.Print("あなたのクラスを教えてね!\n(1:IE、2:SK、3:SE、それ以外:わからない> ")
	var num int
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("4 年制コースなのね!")
	case 2:
		fmt.Println("3 年制コースなのね!")
	case 3:
		fmt.Println("2 年制コースなのね!")
	default:
		fmt.Println("わからないのね!")
	}
}
