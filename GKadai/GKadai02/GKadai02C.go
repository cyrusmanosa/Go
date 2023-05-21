package main

var a, b = 5, 10

func main() {
	println("変数 a の値を表示します!\n", a)
	println("変数 b の値を表示します!\n", b)
	println("変数bにaを代入します!")
	b = a
	println("変数 a の値を表示します!\n", a)
	println("変数 b の値を表示します!\n", b)
}
