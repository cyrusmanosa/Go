package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	YN, name := "", [4]string{"ボロンゴ", "プックル", "チロル", "ゲレゲレ"}
	fmt.Println("わしはこの世界の王、ECC キングじゃ")
	fmt.Println("冒険に出たいという若者はお前か?")
	fmt.Print("名は何という?>")
	fmt.Scan(&YN)
	fmt.Println(YN, "は弱そうじゃ")
	fmt.Println(" ")

	for {
		c, r := 0, rand.Intn(4)
		fmt.Print(name[r], "チロルはどうじゃ?(1:やだ、それ以外:うん)> ")
		fmt.Scan(&c)
		if c != 1 {
			fmt.Printf("そうか!%vを気に入ってもらえたか!\n行け!%v!!世界の平和を取り戻すのじゃ!", name[r], name[r])
			os.Exit(0)
		}
	}
}
