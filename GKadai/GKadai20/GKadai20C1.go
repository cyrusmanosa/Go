package main

func main() {
	for i := 1; i <= 100; i++ {
		print(i)
		if i%2 == 0 {
			print(":偶数です!")
		} else {
			print(":奇数です!")
		}
		println(" ")
	}
}
