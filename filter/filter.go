package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

/*
Filter takes a slice and a callback function as an argument. The callback function returns a Boolean.
Filter passes the slice's elements one by one through the callback function: if it returns true, the
element will be included in the new slice.

In that situation, we had to make use of a type assertion.
*/
func main() {
	baseSlice := []int{1, 2, 3, 4, 5}
	newSlice := funk.Filter(baseSlice, func(x int) bool {
		return x != 2
	}).([]int)

	fmt.Println(baseSlice)
	fmt.Println(newSlice)
}
