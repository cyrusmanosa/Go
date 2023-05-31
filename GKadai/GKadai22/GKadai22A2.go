package main

import "fmt"

func main() {
	for i := 9; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}
