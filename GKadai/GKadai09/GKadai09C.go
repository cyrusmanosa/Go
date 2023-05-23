package main

import "fmt"

func main() {
	var apple, person int
	fmt.Println("ひとり一つずつリンゴを食べます!")
	fmt.Print("リンゴの数を入力してください> ")
	fmt.Scan(&apple)
	fmt.Print("食べる人の数を入力してください> ")
	fmt.Scan(&person)
	if apple < person {
		fmt.Println("リンゴの数が足りません!")
	} else {
		fmt.Println("ひとり一つずつリンゴを食べた!")
		fmt.Println("残ったリンゴは", apple%person, "個です!")
	}

}
