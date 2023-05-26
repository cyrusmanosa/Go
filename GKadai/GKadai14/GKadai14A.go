package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("ここは ECC 苦情処理センターです!")
	fmt.Println("優秀なスタッフが対応します!")
	for {
		r := rand.Intn(100)
		var num int
		fmt.Println("苦情番号:", r, "を受け付けた!")
		fmt.Print("どうしますか?(1:対応する、それ以外:もうやだ)> ")
		fmt.Scan(&num)
		if num == 1 {
			nobita(r)
			fmt.Println(" ")
		} else {
			fmt.Println("おつかれさまでした!")
			os.Exit(0)
		}
	}
}

func nobita(n int) {
	fmt.Println("のび太「ぼくにまかせて!")
	if n < 20 {
		fmt.Println("のび太が対応しました!")
	} else {
		fmt.Println("・・・ジャイアン、お願い!")
		jaian(n)
	}
}

func jaian(n int) {
	fmt.Println("ジャイアン「オレさまにまかせろ!」")
	if n%5 == 0 {
		fmt.Println("ジャイアンが対応しました!")
	} else {
		fmt.Println("・・・スネ夫、お前にゆずってやる!")
		suneo(n)
	}
}

func suneo(n int) {
	fmt.Println("スネ夫「ぼくの出番だ!」")
	if n%3 == 0 {
		fmt.Println("スネ夫が対応しました!")
	} else {
		fmt.Println("スネ夫「・・・出木杉!まかせた!」")
		dekisugi(n)
	}
}

func dekisugi(n int) {
	fmt.Println("出木杉「こんなの簡単さ」")
	fmt.Println("出木杉が対応しました!")
}
