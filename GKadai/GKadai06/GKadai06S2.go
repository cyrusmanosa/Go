package main

import (
	"fmt"
	"math"
)

const PI float64 = 3.14159265

var a int

func main() {
	fmt.Print("半径を入力してください> ")
	fmt.Scan(&a)
	ans := float64(a)

	表面 := 4 * PI * (math.Pow(4, 2))
	体 := (4 * PI * math.Pow(4, 3)) / 3

	fmt.Println(" ")
	fmt.Print("半径", ans, "の球の表面積は", 表面, "、体積は", 体, "です!")
}
