package main

var fuck = map[string]int{
	"ABC": 123,
	"DEF": 456,
}

func main() {
	for k, v := range fuck {
		println(k)
		println(v)
	}
}
