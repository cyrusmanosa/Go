package main

import "fmt"

var C int

func main() {
	println("あなたは勇者です!\nがんばって城までたどり着いてください!!\n")
	gotoTown()
}

func gotoTown() {
	println("あなたは町にいます。")
	print("どちらへ進みますか?(1:草原、2:森)> ")
	fmt.Scan(&C)
	println(" ")
	if C == 1 {
		gotoGrassland()
	} else {
		gotoForest()
	}
	print("←町")
}
func gotoGrassland() {
	println("あなたは草原にいます。")
	print("どちらへ進みますか?(1:町、2:森)> ")
	fmt.Scan(&C)
	println(" ")

	if C == 1 {
		gotoTown()
	} else {
		gotoForest()
	}
	print("←草原")
}
func gotoForest() {
	println("あなたは森にいます。")
	print("どちらへ進みますか?(1:草原、2:荒地)> ")
	fmt.Scan(&C)
	println(" ")

	if C == 1 {
		gotoGrassland()
	} else {
		gotoWasteland()
	}
	print("←森")
}
func gotoWasteland() {
	println("あなたは荒地にいます。")
	print("どちらへ進みますか?(1:草原、2:洞窟)> ")
	fmt.Scan(&C)
	println(" ")

	if C == 1 {
		gotoGrassland()
	} else {
		gotoCave()
	}
	print("←荒地")
}
func gotoCave() {
	println("あなたは洞窟にいます。")
	print("どちらへ進みますか?(1:森、2:城)> ")
	fmt.Scan(&C)
	println(" ")

	if C == 1 {
		gotoForest()
	} else {
		gotoCastle()
	}
	print("←洞窟")
}
func gotoCastle() {
	println("城に着きました!")
	println("履歴を表示します!")
	print("城")
}
