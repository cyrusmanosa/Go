package main

import "fmt"

func main() {
	println("出席状況を入力してください。")
	var 出席, 欠席, 遅刻, 出席率 float32 = 0, 0, 0, 0
	word := " "
	for i := 1; i <= 15; i++ {
		num := 0
		fmt.Print("第", i, "週目:(0:出席、1:欠席、2:遅刻)> ")
		fmt.Scan(&num)
		switch num {
		case 0:
			出席++
		case 1:
			欠席++
		case 2:
			遅刻++
		default:
			i--
		}
	}
	fmt.Printf("出席:%v 回\n欠席:%v 回\n遅刻:%v 回\n", int(出席), int(欠席), int(遅刻))

	if 遅刻 == 3 {
		遅刻 = 0
		欠席++
	}

	出席率 = ((出席 + 遅刻) / 15) * 100

	if 出席率 > 80 {
		word = "OK"
	} else {
		word = "NG"
	}

	fmt.Println("出席率は", 出席率, "%です!出席率", word, "です!")
}
