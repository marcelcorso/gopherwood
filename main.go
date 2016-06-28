package main

import "fmt"

func createNode(key string) (chan string, chan string) {
	in := make(chan string)
	out := make(chan string)

	var (
		rightTo  chan string
		rightFro chan string
		leftTo   chan string
		leftFro  chan string
	)

	go func() {
		for {
			switch <-in {
			// get and set are not used but are didactic
			case "get":
				out <- key
			case "set":
				key = <-in

			case "add":
				newKey := <-in

				sideTo := &rightTo
				sideFro := &rightFro
				if key > newKey {
					sideTo = &leftTo
					sideFro = &leftFro
				} else if key == newKey {
					continue
				}

				if *sideTo == nil {
					*sideTo, *sideFro = createNode(newKey)
				} else {
					*sideTo <- "add"
					*sideTo <- newKey
				}
			case "search":
				searchKey := <-in

				sideTo := &rightTo
				sideFro := &rightFro
				if key > searchKey {
					sideTo = &leftTo
					sideFro = &leftFro
				} else if key == searchKey {
					out <- "found"
					continue
				}

				if *sideTo == nil {
					out <- "nope"
				} else {
					*sideTo <- "search"
					*sideTo <- searchKey
					out <- <-*sideFro
				}
			default:
				out <- "what?"
			}
		}
	}()

	return in, out
}

type Tree struct {
	rootTo  chan string
	rootFro chan string
}

func (t *Tree) Add(key string) {
	if t.rootTo == nil {
		t.rootTo, t.rootFro = createNode(key)
	} else {
		t.rootTo <- "add"
		t.rootTo <- key
	}
}

func (t *Tree) Search(key string) (response string) {
	if t.rootTo != nil {
		t.rootTo <- "search"
		t.rootTo <- key
		response = <-t.rootFro
	}
	return
}

type Node struct {
	key   string
	left  *Node
	right *Node
}

func main() {
	t := Tree{}
	keys := []string{"ma", "ro", "zin", "ko", "aq", "er", "se", "ca", "pi", "ty", "ge", "me", "mo"}
	for i := 0; i < len(keys); i++ {
		t.Add(keys[i])
	}
	fmt.Println("")

	found := t.Search("se")
	fmt.Printf("'se' is in the tree? %v\n", found)
}
