package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("ひとつめの整数を入力してください> ")
	fmt.Scan(&num1)
	fmt.Print("ふたつめの整数を入力してください> ")
	fmt.Scan(&num2)
	fmt.Printf("最大公約数は%vです!\n", calcGCD(num1, num2))
	fmt.Printf("最小公倍数は%vです!", num1*num2/calcLCM(num1, num2))
}

func calcGCD(n1, n2 int) int {
	ans1 := 0
	for i := 2; i < n1; i++ {
		if n1%i == 0 && n2%i == 0 {
			ans1 = i
		}
	}
	return ans1
}

func calcLCM(n1, n2 int) int {
	if n2 == 0 {
		return n1
	}
	return calcLCM(n2, n1%n2)
}
