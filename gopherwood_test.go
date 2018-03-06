package gopherwood

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func TestNothing(t *testing.T) {}

func TestSearch(t *testing.T) {
	rand.Seed(boringRandomSeed)

	noKeys := 100
	keySize := 30
	keys := make([]string, noKeys)

	tree := Tree{}

	for n := 0; n < noKeys; n++ {
		k := randStringBytesRmndr(keySize)
		keys[n] = k
		tree.Add(k)
	}

	for n := 0; n < noKeys; n++ {
		k := keys[n]
		got := tree.Search(k)
		if k != got {
			t.Error("want", k, "but got", got)
		}
	}
}

func TestConcurrentSearch1pct(t *testing.T) {
	rand.Seed(boringRandomSeed)

	noKeys := 15
	keySize := 30
	keys := make([]string, noKeys)

	tree := Tree{}

	for n := 0; n < noKeys; n++ {
		k := randStringBytesRmndr(keySize)
		keys[n] = k
		tree.Add(k)
	}

	//noSearchers := 1 * noKeys / 100
	noSearchers := 2
	var wg sync.WaitGroup
	wg.Add(noSearchers)

	sliceSize := noKeys / noSearchers

	fmt.Println("noKeys: ", noKeys)
	fmt.Println("noSearchers: ", noSearchers)
	fmt.Println("sliceSize: ", sliceSize)

	for n := 0; n < noSearchers; n++ {
		go func(n int) {
			base := n * sliceSize
			for m := 0; m < sliceSize; m++ {
				fmt.Println("n: ", n, "base: ", base, "base+m: ", base+m)
				key := keys[base+m]
				fmt.Println(key)
				got := tree.Search(key)
				if key != got {
					t.Error("want", key, "but got", got)
				}
			}
			wg.Done()
		}(n)
	}
	wg.Wait()
	fmt.Println("wow")
}

func BenchmarkGopherwoodAdd(b *testing.B) {
	rand.Seed(boringRandomSeed)
	t := Tree{}
	for n := 0; n < b.N; n++ {
		t.Add(randStringBytesRmndr(b.N))
	}
}
