package main

import "math/rand"

func main() {
	cards := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	dst := 0
	for dst < len(cards)-1 {
		rest := len(cards) - dst
		src := dst + rand.Intn(rest)
		tmp := cards[dst]
		if dst == src {
			println("同じカードなので入れ替えをパスします!")
			continue
		}
		cards[dst] = cards[src]
		cards[src] = tmp
		println(dst, "番目と", src, "番目を入れ替えました!")
		dst++
	}
	for i := 0; i < len(cards); i++ {
		print(cards[i], " ")
	}
	println(" ")
}
