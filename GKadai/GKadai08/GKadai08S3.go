package main

import "fmt"

func main() {
	var nh, nm, th, tm, ah, am int
	fmt.Print("現在の時刻(時)を入力してください(0-23)> ")
	fmt.Scan(&nh)
	fmt.Println(" ")
	fmt.Print("在の時刻(分)を入力してください(0-59)> ")
	fmt.Scan(&nm)
	fmt.Println(" ")

	fmt.Print("授業終了の時刻(時)を入力してください(0-23)> ")
	fmt.Scan(&th)
	fmt.Println(" ")
	fmt.Print("授業終了の時刻(分)を入力してください(0-59)> ")
	fmt.Scan(&tm)

	if th == 0 {
	}

	if tm > nm {
		am = tm - nm
		ah = th - nh
	} else {
		am = (tm + 60) - nm
		ah = (th - 1) - nh
	}

	if ah < 0 {
		fmt.Println("君の人生も終わった")
	} else {
		fmt.Println("あと", ah, "時間", am, "分、がんばれ!")
	}
}
