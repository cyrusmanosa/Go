package main

import "fmt"

func main() {
	var num1, num2, People int
	fmt.Print("1 より大きい整数 1 を入力してください> ")
	fmt.Scan(&num1)
	fmt.Print("1 より大きい整数 2 を入力してください> ")
	fmt.Scan(&num2)
	fmt.Printf("%v÷%vの商と余りを計算します!", num1, num2)
	fmt.Print("誰を呼び出しますか?(1:出木杉、それ以外:のび太)> ")
	fmt.Scan(&People)
	if People == 1 {
		dekisugi(num1, num2)
	} else {
		fmt.Println("あ、そうか!リンゴを食べる問題と一緒だね!")
		nobita(num1, num2)
	}
}

func dekisugi(n1, n2 int) {
	fmt.Println("出木杉くん、よろしく!")
	fmt.Println("そんなの簡単さ!")
	a := n1 / n2
	b := n1 % n2
	fmt.Println("商は", a, "、余りは", b, "だよ。")
}

func nobita(n1, n2 int) {
	var ans1, ans2, r int
	if n1 > n2 {
		for n1 > n2 {
			n1 -= n2
			r++
		}
		ans1 = r
		ans2 = n1
	} else {
		for n2 > n1 {
			n2 -= n1
		}
		ans1 = n2
		ans2 = n1
	}
	fmt.Printf("商は%v、余りは%vだよ。", ans1, ans2)
}
