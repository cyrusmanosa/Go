package main

func main() {
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 1; j++ {
			println("********")
		}
		if i == 2 {
			break
		}
		println("*")
	}
}
