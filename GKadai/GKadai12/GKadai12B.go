package main

import "fmt"

func main() {
	people, mask := 1, 0
	array := [...]int{9: 0}
	for {
		fmt.Print(people, "人目の点数を入力してください> ")
		fmt.Scan(&array[people])
		if array[people] < 0 {
			people--
			break
		}
		mask += array[people]
		people++
	}

	fmt.Println(people, "人の合計は", mask, "点です!")
	fmt.Println(people, "人の平均は", float32(mask)/float32(people), "点です!")

}
