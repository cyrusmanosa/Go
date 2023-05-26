package main

func main() {
	num := [5]int{30, 50, 100, 70, 95}
	println("最高得点を探します!")
	ans := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > ans {
			ans = num[i]
		}
	}
	println("最高得点は", ans, "点です!")
}
