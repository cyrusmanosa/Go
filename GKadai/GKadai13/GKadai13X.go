package main

import "fmt"

func main() {
	var num int
	fmt.Print("いくつまで叫びますか?> ")
	fmt.Scan(&num)
	isFizz(num)
}

func isFizz(num int) {
	for i := 1; i <= num; i++ {

		switch {
		case i%15 == 0:
			fmt.Println("わにゃん!!")
		case i%3 == 0 || i/10 == 3:
			fmt.Println("わん!")
		case i%5 == 0 || i/10 == 5:
			fmt.Println("にゃん!")
		default:
			fmt.Println(i)
		}
	}
}
