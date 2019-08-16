package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

/*
The Reduce function takes the following arguments: a slice, a callback function and, an initial value for the accumulator.
In Go-Funk, this function returns a float.

With the Reduce function, the type assertion is not necessary. If you have a look at the function signature, you'll notice
that it will always return a float.
*/
func main() {
	baseSlice := []int{1, 2, 3, 4, 5}
	sum := funk.Reduce(baseSlice, func(acc, elem int) int { return acc + elem }, 0)

	fmt.Println(baseSlice)
	fmt.Println(sum)
}
