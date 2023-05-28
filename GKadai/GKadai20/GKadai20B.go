package main

import (
	"fmt"
	"math/rand"
)

func main() {
	card := [9]int{}
	for i := 0; i < len(card); i++ {
		r := rand.Intn(100)
		card[i] = r
		fmt.Printf("card[%v]:%v\n", i, card[i])
	}
	fmt.Println("逆順で表示します!")
	for i := len(card) - 1; i >= 0; i-- {
		fmt.Printf("card[%v]:%v\n", i, card[i])
	}
}
