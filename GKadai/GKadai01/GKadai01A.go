package main

func main() {
	E()
	C()
	C()
}

func E() {
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 1; j++ {
			println(" ******")
		}
		if i == 2 {
			break
		}
		println(" *")
	}
}

func C() {
	var k int
	for k = 0; k < 7; k++ {
		switch k {
		case 0, 6:
			println("  **** ")
			break
		case 1, 5:
			println(" *    *")
			break
		default:
			println("*")
			break
		}
	}
}
