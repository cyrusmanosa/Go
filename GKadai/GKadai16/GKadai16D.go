package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	for i := 0; i < len(a); i++ {
		fmt.Println("a[", i, "]:", a[i])
	}
	fmt.Println("a[0]と a[2]を入れ替えます!")
	for i := 0; i < len(a); i++ {
		fmt.Println("a[", i, "]:", a[len(a)-1-i])
	}
}
