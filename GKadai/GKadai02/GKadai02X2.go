package main

func main() {
	var n = 100
	println("nの値\n", n)

	for i := 99; i > 0; i-- {
		n += i
	}
	println("*** 1からnまで加算します! ***\n", "結果は\n", n)
}
