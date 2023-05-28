package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

func main() {
	num, r, Yans := 0, 1, 0
	fmt.Print("生成する整数の桁数を入力してください(1~9)> ")
	fmt.Scan(&num)
	x := int(math.Pow10(num))
	ans := rand.Intn(x)

	fmt.Println("----", ans, "-------")

	for {
		fmt.Printf("4 桁の整数を入力してください[%v]> ", r)
		fmt.Scan(&Yans)

		if Yans == ans {
			fmt.Printf("\n%v......正解!", r)
			os.Exit(0)
		} else if Yans/ans < 0 {
			r--
			continue
		}

		point := 0
		Cans := ans

		for i := 1; i <= num; i++ {
			check := Cans % 10
			fmt.Println("i=", i, "=====", check)
			CYans := Yans
			for j := 1; j <= num; j++ {
				cya := CYans % 10
				fmt.Println("\tj=", j, "-----", cya)

				if check == cya {
					point++
				}
				if check == cya && j == i {
					point += 10
				}
				CYans /= 10
			}
			Cans /= 10
		}

		fmt.Println(Yans, "......点数は", point)
		r++
	}
}
