package main

import "fmt"

func main() {
	var array [10]int
	for i := 1; i < 10; i++ {
		if i == 1 || i == 2 {
			array[i] = 1
		} else {
			array[i] = array[i-1] + array[i-2]
		}

		fmt.Println("第", i, "項:\t", array[i])
	}
}
