package main

import (
	"fmt"
	"simple-go-hash-table/lib"
)

func main() {
	hash := lib.HashData[string]{}
	hash.Put("foo", "bar")
	hash.Put("fof", "bar_2")

	get1 := hash.Get("fof")
	fmt.Println(get1) // bar_2

	get2 := hash.Get("foo")
	fmt.Println(get2) // bar

	hash.Remove("fof")
	get3 := hash.Get("fof")
	fmt.Println(get3) // empty
}
