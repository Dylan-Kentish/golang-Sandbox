package main

import (
	"fmt"
	"strings"
)

// Transform applies the transformation to every item in the slice m and returns the result.
func Transform[S []T, T any, U any](m S, transformation func(item T) U) []U {
	r := make([]U, 0, len(m))
	for _, v := range m {
		r = append(r, transformation(v))
	}
	return r
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	transformed := Transform(data, func(item int) int {
		return -item 
	})

	fmt.Println(data)
	fmt.Println(transformed)

	data2 := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	transformed2 := Transform(data2, func(item string) string {
		return strings.ToLower(item)
	})

	fmt.Println(data2)
	fmt.Println(transformed2)
}
