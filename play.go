package main

import "fmt"

func createNode(key string) (chan string, chan string) {
	in := make(chan string)
	out := make(chan string)

	go func() {
		for {
			method := <-in
			if method == "get" {
				out <- key
			} else if method == "set" {
				key = <-in
				out <- "ok"
			} else {
				out <- "what?"
			}
		}
	}()

	return in, out
}

func main() {

	fmt.Println("Hello, playground")

	to, fro := createNode("zarabum")
	to <- "get"
	fmt.Println(<-fro)

	to <- "set"
	to <- "monster"
	fmt.Println(<-fro)

	to <- "get"
	fmt.Println(<-fro)

}
