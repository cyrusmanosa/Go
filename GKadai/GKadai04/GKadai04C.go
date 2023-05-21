package main

import (
	"fmt"
)

func main() {
	var name string
	var ago int
	fmt.Print("あなたのお名前は?>")
	fmt.Scan(&name)
	fmt.Println(name, "さん、こんにちは!")
	fmt.Print("年齢はいくつ?>")
	fmt.Scan(&ago)
	fmt.Println(ago, "歳なのね!")
}
