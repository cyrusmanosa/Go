package main

func main() {
	menu := map[string]int{
		"コーヒー":    290,
		"さくらドーナツ": 250,
		"チキンサラダ":  380,
		"チーズケーキ":  420,
	}
	rate := 1.1
	println("ようこそ!ECC コーヒーへ", "\n", "こちらでお召し上がりですね!\n", "--------\n")

	for k, v := range menu {
		println(k, "\t", v, "円\t小計：", v*int(rate), "円")
	}
}
