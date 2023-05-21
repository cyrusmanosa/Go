package main

var x = 5

func main() {
	println("変数 x の値を表示します!\n", x)
	x += 14
	println("変数xに14を足します!\n", x)
	x -= 3
	println("変数xから3を引きます!\n", x)
	x *= 7
	println("変数xに7をかけます!\n", x)
	x /= 10
	println("変数xを10で割ります!\n", x)
	x %= 4
	println("変数 x を 4 で割った余りを求めます!\n", x)
}
