package main

func main() {
	println("はじめてのおつかい!\nのび太がケーキを買いに行きます!")
	gotoPark()
	gotoPostOffice()
	gotoBank()
	gotoCakeShop()
	gotoBank()
	gotoPostOffice()
	gotoPark()
	println("無事戻ってきました!")
}

func gotoPark() {
	println("公園に着きました!")
}

func gotoPostOffice() {
	println("郵便局に着きました!")
}

func gotoBank() {
	println("銀行に着きました!")
}
func gotoCakeShop() {
	println("ケーキ屋に着きました!\n*** ケーキを買いました! ***\n戻ります!")
}
