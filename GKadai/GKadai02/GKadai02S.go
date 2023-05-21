package main

func main() {
	var x, y = 2, 7
	var z = y
	println("x の値は\n", x, "\ny の値は\n", y, "\n*** x と y の値を入れ替えます! ***\n", z)
	z = x
	println("y の値は\n", z)
}
