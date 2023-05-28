package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	num, choose := 0, [5]string{"騎士", "弓兵", "槍兵", "将軍", "スパイ"}
	for {
		cpu := rand.Intn(5)
		fmt.Println("あなたは何を選びますか?")
		fmt.Print(" ( ")
		for i := 0; i < len(choose); i++ {
			fmt.Print(i, ":", choose[i])
		}
		fmt.Print(" ) > ")
		fmt.Scan(&num)
		if num < 0 {
			os.Exit(0)
		}
		fmt.Println("あなたは", choose[num], "を出した!")
		fmt.Println("CPUは", choose[cpu], "を出した!")
		switch {
		case num == cpu:
			fmt.Println("あいこだ!")
		case num == 0 && cpu == 1:
			fmt.Println("あなたの勝ち!")
		case num == 1 && cpu == 2:
			fmt.Println("あなたの勝ち!")
		case num == 2 && cpu == 0:
			fmt.Println("あなたの勝ち!")
		case num == 3 && cpu == 4:
			fmt.Println("あなたの勝ち!")
		case num == 4 && cpu == 3:
			fmt.Println("あなたの勝ち!")
		default:
			fmt.Println("あなたの負け!")
		}
		fmt.Println(" ")
	}
}
