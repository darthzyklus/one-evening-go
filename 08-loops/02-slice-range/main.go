package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(numbers ...int) int {
	result := 0

	for _, n := range numbers {
		result += n
	}

	return result
}
