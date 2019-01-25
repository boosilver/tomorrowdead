package main

import (
	"fmt"
)

func main() {
	// values := []interface{}{"mydata", 1234567890.123}
	g := -3332555.
	values := []interface{}{"ans", g}
	for _, g := range values {
		// fmt.Printf("%.4e\n", v)
		fmt.Printf("%.5e\n", g)
	}
}
