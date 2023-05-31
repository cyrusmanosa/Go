package main

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j < i; j++ {
			print(" ")
		}
		for j := 9; j >= i; j-- {
			print("*")
		}
		if i != 9 {
			println(" ")
		}
	}
	for i := 0; i <= 9; i++ {
		for j := 9; j > i; j-- {
			print(" ")
		}
		for j := 1; j <= i; j++ {
			print("*")
		}
		println(" ")
	}
}
