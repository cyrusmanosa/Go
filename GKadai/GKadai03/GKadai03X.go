package main

import "fmt"

func printMulTable(k int) {
	for i := 1; i <= k; i++ {
		for j := 1; j <= k; j++ {
			print("\t", j*i)
		}
		println()
	}
}

func main() {
	var text int
	print("Enter number: ")
	fmt.Scan(&text)
	printMulTable(text)
}
