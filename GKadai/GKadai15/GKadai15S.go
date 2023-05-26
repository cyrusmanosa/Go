package main

import "fmt"

func main() {
	var Team, People int
	fmt.Print("トーナメントのチーム数を入力してください> ")
	fmt.Scan(&Team)
	fmt.Print("誰が試合数を計算しますか?(1:出木杉、それ以外:しずか)> ")
	fmt.Scan(&People)
	fmt.Println("そんなの簡単さ!")

	if People == 1 {
		fmt.Println("トータルの試合数は", dekisugi(Team), "だよ。")
	} else {
		fmt.Println("最後に 1 チーム残るためには", sizuka(Team), "試合すればいいのよ。")
	}
}

func sizuka(T int) int {
	fmt.Println("1 試合すれば●●●●●●●●●●")
	return T - (T % 2)
}

func dekisugi(T int) int {
	i := 1
	var r, nB, allR int

	for T > 1 {
		r = T / 2
		nB = T % 2
		T = r + nB
		fmt.Printf("%v回戦の試合数:%v\t不戦勝:%v \t残りチーム:%v\n", i, r, nB, T)
		i++
		allR += r
	}
	return allR
}
