package main

import (
	"fmt"
)

// Slices are a key data type in Go
func main() {
	// To create an empty slice with non-zero length, use the builtin make
	s := make([]string, 3) // make a slice of strings of length 3
	fmt.Println(s)         // []

	// can set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s[2]) // c
	fmt.Println(s)    // [a b c]
	// "len" returns the length of the slice
	fmt.Println(len(s)) // 3

	// builtin "append" , which returns a slice containing one or more new values
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]

	// create an empty slice c of the same length as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) // [a b c d e f]

	// Slices support a “slice” operator with the syntax "slice[low:high]"
	l := s[2:5]    // this gets a slice of the elements s[2], s[3], and s[4]
	fmt.Println(l) // [c d e]

	m := s[:5]     // This slices up to (but excluding) s[5]
	fmt.Println(m) // [a b c d e]

	n := s[2:]     // And this slices up from (and including) s[2]
	fmt.Println(n) // [c d e f]

	// can declare and initialize a variable for slice in a single line as well
	o := []string{"g", "h", "i"}
	fmt.Println(o) // [g h i]

	// Slices can be composed into multi-dimensional data structures
	// The length of the inner slices can vary, unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD) // [[0] [1 2] [2 3 4]]

}
