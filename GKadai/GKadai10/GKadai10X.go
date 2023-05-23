package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var you int
	array := [3]string{"グー", "チョキ", "パー"}
	fmt.Println("じゃんけんをします!")
	fmt.Print("何の手を出しますか?(0:グー、1:チョキ、2:パー> ")
	fmt.Scan(&you)
	if you > 2 || you < 0 {
		os.Exit(0)
	}
	fmt.Println("あなたは", array[you], "を出した!")
	comp := rand.Intn(3)
	fmt.Println(comp)
	fmt.Println("コンピュータは", comp, "を出した!")
	switch {
	case you == comp:
		fmt.Println("あいこだ!")
	case (you == 0 && comp == 2) || (you == 1 && comp == 0) || (you == 2 && comp == 1):
		fmt.Println("あなたの負け!")
	default:
		fmt.Println("あなたの勝ち!")
	}
}
