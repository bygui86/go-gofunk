package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

/*
The Map function iterates through the elements of a slice (or a map) and modifies them. It returns a new array.

The result is what we could expect from this code. But if you have a closer look, you’ll see an issue with the type of
the newSlice variable: it is interface{}. To fix this problem, we'll make use of type assertions. In order to do this,
you'd just have to add .([]int) right after funk.Map(...). This way, we're telling the Go compiler: "You can't determine
what's the type of this value, but I assure you it's a slice of int, so could you please convert it ?".
*/
func main() {

	generic()

	mapToSlice()

	mapToMap()

	sliceToMap()

	sliceToSlice()
}

func generic() {

	baseSlice := []int{1, 2, 3, 4, 5}
	newSlice := funk.Map(baseSlice, func(x int) int {
		return x + 1
	}).([]int)

	fmt.Println(baseSlice)
	fmt.Println(newSlice)
}

func mapToSlice() {

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	funk.Map(mapping, func(k int, v string) int {
		return k
	}) // []int{1, 2}
}

func mapToMap() {

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	funk.Map(mapping, func(k int, v string) (string, string) {
		return fmt.Sprintf("%d", k), v
	}) // map[string]string{"1": "Florent", "2": "Gilles"}
}

func sliceToMap() {

	funk.Map([]int{1, 2, 3, 4}, func(x int) (int, int) {
		return x, x
	}) // map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
}

func sliceToSlice() {

	funk.Map([]int{1, 2, 3, 4}, func(x int) int {
		return x * 2
	}) // []int{2, 4, 6, 8}

	funk.Map([]int{1, 2, 3, 4}, func(x int) string {
		return "Hello"
	}) // []string{"Hello", "Hello", "Hello", "Hello"}
}
