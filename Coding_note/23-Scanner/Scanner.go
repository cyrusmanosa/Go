package main

import (
	"bufio"
	"fmt"
	"os"
)

func V1() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(text)
		} else {
			break
		}
	}
	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}

func V2() {

	var val int
	// 数値の入力
	fmt.Println("数値を入力してください")
	fmt.Scan(&val)

	fmt.Println("入力結果")
	fmt.Printf("%d\n", val)

	// 文字列の入力
	fmt.Println("文字列を入力してください")
	var str string
	fmt.Scan(&str)

	fmt.Println("入力結果")
	fmt.Println(str)
}

func main() {
	V2()
}
