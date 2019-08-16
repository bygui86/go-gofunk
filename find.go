package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

func main() {

	r := funk.Find([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	}) // 2
	fmt.Println("Result:", r)
}
