package main

import (
	"fmt"
	"example.com/user/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes(".dlrow ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}