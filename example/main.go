package main

import (
	"fmt"

	"github.com/marcelcorso/gopherwood"
)

func main() {
	t := gopherwood.Tree{}
	keys := []string{"ma", "ro", "zin", "ko", "aq", "er", "se", "ca", "pi", "ty", "ge", "me", "mo"}
	for i := 0; i < len(keys); i++ {
		t.Add(keys[i])
	}
	fmt.Println("")

	found := t.Search("se")
	fmt.Printf("'se' is in the tree? %v\n", found)
}
