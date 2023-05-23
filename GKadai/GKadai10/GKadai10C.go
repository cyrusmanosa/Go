package main

import "fmt"

func main() {
	var num int
	fmt.Println("こんにちは!")
	fmt.Println("わたしはスーパーティーチャー、ECC エクセレントよ!")
	fmt.Print("あなたのコースを教えてね!(1:IE、2:SK、3:SE)> ")
	fmt.Scan(&num)
	if num > 3 {
		fmt.Println("ちゃんと答えてね!")
	} else if num > 0 && num < 4 {
		fmt.Println("IT カレッジの学生ね!")
	}
}
