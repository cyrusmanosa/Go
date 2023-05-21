package main

import (
	"fmt"
	"math"
)

func main() {
	const PI = 3.14
	var R float32

	fmt.Print("半径を入力してください>")
	fmt.Scan(&R)

	表面 := 4 * PI * (math.Pow(4, 2))
	体 := (4 * PI * math.Pow(4, 3)) / 3
	fmt.Print("\n半径")
	fmt.Printf("%v の球の表面積は %v、体積は %v です!", R, 表面, 体)
}
