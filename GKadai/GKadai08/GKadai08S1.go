package main

import "fmt"

func main() {
	var speed, m float64
	fmt.Println("A さんの前を列車が通過します!")
	fmt.Print("列車の速度(時速)を入力してください> ")
	fmt.Scan(&speed)
	fmt.Println(" ")
	fmt.Print("列車の長さ(メートル)を入力してください> ")
	fmt.Scan(&m)
	s := (m * 1000 / speed * 3600)
	fmt.Println(s, "秒かかります!")
}
