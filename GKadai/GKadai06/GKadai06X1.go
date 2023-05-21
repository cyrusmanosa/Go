package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var name string
	var age int
	fmt.Println("こんにちは!")
	fmt.Println("わたしは占いマシーンの ECC1000 よ!")
	fmt.Println("あなたのことを占ってあげるわ!よろしくね")

	fmt.Println(" ")
	fmt.Print("名前は何ていうの?> ")
	fmt.Scan(&name, "\n")

	fmt.Print("年齢はいくつ?> ")
	fmt.Scan(&age, "\n")
	fmt.Println(" ")

	a := rand.Intn(100)
	b := rand.Intn(100)
	c := rand.Intn(100)

	fmt.Println(name, "さん、こんにちは!\nあなたは", age, "歳なんですね!")
	fmt.Println(" ")
	fmt.Println(name, "さんの今日の運勢は")
	fmt.Println("ラブ運\t", a)
	fmt.Println("金銭運\t", b)
	fmt.Println("全体運\t", c)
	fmt.Println(" ")
	fmt.Println("また来てね!")
}
