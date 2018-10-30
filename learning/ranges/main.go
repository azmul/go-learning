package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	// use range to sum the numbers in a slice
	for _, num := range nums { // ignore index with the blank identifier _
		sum += num
	}
	fmt.Println(sum) // 6

	for i, num := range nums { // slices provides both the index and value for each entry
		if num == 3 {
			fmt.Println(i) // 2
		}
	}

	// range on map iterates over key/value pairs.
	keyValues := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range keyValues {
		fmt.Println(key, value) // a apple   b banana
	}

	// range can also iterate over just the keys of a map.
	for key := range keyValues {
		fmt.Println(key) // a  b
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune
	// and the second the rune itself
	for i, c := range "abc" {
		fmt.Println(i, c) // 0 97   1 98   2 99
	}
}
