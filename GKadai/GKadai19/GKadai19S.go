package main

import (
	"fmt"
	"math/rand"
)

func main() {
	name, stone, c, cpu := "", 20, 0, rand.Intn(3)+1
	fmt.Print("あなたの名前を入力してください> ")
	fmt.Scan(&name)
	fmt.Println(" ")
	for {
		showStone(stone)
		fmt.Println("\n", name, "の番です。")
		fmt.Print("何個取りますか?(1-3)> ")
		fmt.Scan(&c)
		if c > 3 {
			continue
		}
		fmt.Println(c, "個取りました!")
		stone -= c
		fmt.Println(" ")
		showStone(stone)
		stone -= cpu
		if stone-cpu == 0 {
			fmt.Println("CPUの勝ちです。")
		} else {
			fmt.Println("CPUの負けです!")
		}
		fmt.Println("\nCPU の番です。")
		fmt.Println(cpu, "個取りました!")
		fmt.Println(" ")

	}
}

func showStone(stone int) int {
	fmt.Print("残り", stone, "個:")
	for i := 0; i < stone; i++ {
		fmt.Print("●")
	}
	return stone
}
