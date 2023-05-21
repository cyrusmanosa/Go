package main

var j = 0

func main() {

	for i := 1; i <= 10; i++ {
		j += i
	}
	print("1 から 10 まで加算します!\n", "結果は\n", j)

}
