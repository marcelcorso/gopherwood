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
			case "get":
				out <- key
			case "set":
				key = <-in
			case "add":
				newKey := <-in
				// TODO pointer to chan can dedup this:
				if key < newKey {
					if rightTo == nil {
						rightTo, rightFro = createNode(newKey)
					} else {
						rightTo <- "add"
						rightTo <- newKey
					}
				} else if key > newKey {
					if rightTo == nil {
						leftTo, leftFro = createNode(newKey)
					} else {
						leftTo <- "add"
						leftTo <- newKey
					}
				}
				// nothing to do if key == newKey
			case "search":
				searchKey := <-in
				if key < searchKey {
					if rightTo == nil {
						out <- "nope"
					} else {
						rightTo <- "search"
						rightTo <- searchKey
						out <- <-rightFro
					}
				} else if key > searchKey {
					if leftTo == nil {
						out <- "nope"
					} else {
						leftTo <- "search"
						leftTo <- searchKey
						out <- <-leftFro
					}
				} else {
					// searchKey == key
					out <- "found"
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
		fmt.Print(keys[i])
		fmt.Print(", ")
	}
	fmt.Println("")

	found := t.Search("se")
	fmt.Printf("'se' is in the tree? %v", found)
}
