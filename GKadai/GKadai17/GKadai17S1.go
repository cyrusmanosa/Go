package main

func main() {
	fibonacci(1)
}

func fibonacci(n int) {
	array := [20]int{}
	array[0], array[1] = 0, n
	for i := 2; i < len(array); i++ {
		array[i] = array[i-1] + array[i-2]
	}
	for i := 0; i < len(array); i++ {
		print(array[i], " ")
	}

}
