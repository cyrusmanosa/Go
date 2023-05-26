package main

import "fmt"

func main() {
	println("ジャックは豆をまいた!")
	println("次の日、天まで届く豆の木に成長した!")
	println("ここを登ればきっと宝物があるに違いない!!")
	claimBeanTree()
	println("めでたし、めでたし。")
}

func claimBeanTree() {
	var num int
	r := 1
	for {
		println("ジャックは豆の木を登った!")
		print("どうしますか?(1:登る、それ以外:もう疲れた)> ")
		fmt.Scan(&num)
		if num == 1 {
			r++
		} else if num >= 2 {
			break
		}
	}
	for r != 0 {
		fmt.Println("ジャックは豆の木を降りた!")
		r--
	}
}
