package main

import (
	"fmt"
)

// Maps are Go’s built-in associative data type
// (sometimes called hashes or dicts in other languages)
func main() {
	// To create an empty map, use the builtin
	// make: make(map[key-type]val-type)
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m)       // map[k1:1 k2:2]
	fmt.Println(m["k2"]) // 2

	// builtin len returns the number of key/value pairs
	fmt.Println(len(m)) // 2

	// builtin delete removes key/value pairs
	delete(m, "k2")
	fmt.Println(m) // map[k1:1]

	_, prs := m["k2"] // blank identifier _
	fmt.Println(prs)  // false

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
