package main

import "fmt"

func main() {
	num := 0
	fmt.Print("1 以上の整数を入力してください> ")
	fmt.Scan(&num)

	sumNumber(num)

	if isMultiple3(num) {
		println("3 の倍数です!")
	} else {
		println("3 の倍数ではありません!")
	}
}

func sumNumber(n int) {
	a, b := 0, 0
	for n/10 == 0 {
		a = n % 10
		n = n / 10
		b = b + a
	}
	b += n
	if b%3 == 0 {
		println("3 の倍数です!")
	} else {
		println("3 の倍数ではありません!")
	}
}

func isMultiple3(n int) bool {
	if n%3 == 0 {
		return true
	} else {
		return false
	}
}
