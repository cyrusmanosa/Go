package main

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			print("\t", j*i)
		}
		println()
	}
}
