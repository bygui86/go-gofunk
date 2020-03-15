package main

import "github.com/thoas/go-funk"

type Foo struct {
	ID        int
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
}

func (f Foo) TableName() string {
	return "foo"
}

// see also, typesafe implementations: ContainsInt, ContainsInt64, ContainsFloat32, ContainsFloat64, ContainsString
func main() {

	stringSamples()

	mapSamples()

	customSamples()
}

func stringSamples() {
	// slice of string
	funk.Contains([]string{"foo", "bar"}, "bar") // true

	// string
	funk.Contains("florent", "rent") // true
	funk.Contains("florent", "foo")  // false
}

func customSamples() {

	// slice of Foo ptr
	f := &Foo{
		ID:        1,
		FirstName: "Foo",
		LastName:  "Bar",
		Age:       30,
	}
	funk.Contains([]*Foo{f}, f)   // true
	funk.Contains([]*Foo{f}, nil) // false

	b := &Foo{
		ID:        2,
		FirstName: "Florent",
		LastName:  "Messa",
		Age:       28,
	}

	funk.Contains([]*Foo{f}, b) // false
}

func mapSamples() {

	funk.Contains(map[int]string{1: "Florent"}, 1) // true
}
