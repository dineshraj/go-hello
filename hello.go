package main

import (
	"fmt"
	"github.com/dineshraj/go-hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
