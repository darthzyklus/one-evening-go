package main

import "fmt"

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}

func Keys(products map[int]string) []int {
	result := []int{}

	for key := range products {
		result = append(result, key)
	}

	return result
}

func Values(products map[int]string) []string {
	result := []string{}

	for _, value := range products {
		result = append(result, value)
	}

	return result

}
