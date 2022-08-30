package main

import "fmt"

var allocations int

func AllocateBuffer() *string {
	if allocations == 3 {
		return nil
	}

	allocations++
	return new(string)
}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
