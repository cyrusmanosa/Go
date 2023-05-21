package main

const COFFEE = 290

var tax, total = 1.08, COFFEE * tax

func main() {
	print("ようこそ!ECC コーヒーへ",
		"\nお持ち帰りですね!",
		"\nコーヒー", COFFEE, " 円、消費税が", (int)(COFFEE*(tax-1)), "円で合計", (int)(total), "円になります! ", "\nありがとうございました!")
}
