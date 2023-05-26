package main

import (
	"fmt"
	"os"
)

func main() {
	var num [1000]int
	var ans int
	for i := 1; i < len(num); i++ {
		fmt.Print(i, "人目の点数を入力してください> ")
		fmt.Scan(&num[i])

		if num[i] < 0 && i == 1 {
			fmt.Println("少なくとも一人分は入力してください!")
			i--
			continue
		}

		if num[i] < 0 && i != 1 {
			fmt.Println(i-1, "人の合計は", ans, "点です!")
			fmt.Println(i-1, "人の平均は", float32(ans)/float32(i-1), "点です!")
			os.Exit(0)
		}
		ans += num[i]
	}

}
