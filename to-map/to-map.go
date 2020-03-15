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

func main() {

	f := &Foo{
		ID:        1,
		FirstName: "Gilles",
		LastName:  "Fabio",
		Age:       70,
	}

	b := &Foo{
		ID:        2,
		FirstName: "Florent",
		LastName:  "Messa",
		Age:       80,
	}

	results := []*Foo{f, b}

	funk.ToMap(results, "ID") // map[int]*Foo{1: f, 2: b}
}
