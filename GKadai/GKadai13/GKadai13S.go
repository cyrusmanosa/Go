package main

import (
	"fmt"
	"os"
)

func main() {
	var num int
	fmt.Println("わたしは量子コンピュータ ECC トロピカルよ!")

	for {
		fmt.Print("何の段を教えてほしいの?> ")
		fmt.Scan(&num)
		if num > 9 {
			fmt.Println("そんな難しいの、わかんないわ!")
			continue
		} else if num < 0 {
			fmt.Println("さよなら!")
			os.Exit(0)
		}
		for i := 1; i < 10; i++ {
			fmt.Print("\t", num*i)
		}
		fmt.Println(" ")
	}
}
