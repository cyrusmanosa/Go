package main

func main() {
	for i := 0; i < 10; i++ {
		println("i = ", i)
		for j := 0; j < 10; j++ {
			print("j = ", j, ",")
		}
		println(" ")
	}
}
