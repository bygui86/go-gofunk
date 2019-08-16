package main

import "github.com/thoas/go-funk"

// see also, typesafe implementations: IndexOfInt, IndexOfInt64, IndexOfFloat32, IndexOfFloat64, IndexOfString
// see also, typesafe implementations: LastIndexOfInt, LastIndexOfInt64, LastIndexOfFloat32, LastIndexOfFloat64, LastIndexOfString
func main() {

	// slice of string
	funk.IndexOf([]string{"foo", "bar"}, "bar")    // 1
	funk.IndexOf([]string{"foo", "bar"}, "gilles") // -1

	// slice of string
	funk.LastIndexOf([]string{"foo", "bar", "bar"}, "bar") // 2
	funk.LastIndexOf([]string{"foo", "bar"}, "gilles")     // -1
}
