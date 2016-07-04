package gopherwood

import (
	"math/rand"
	"testing"
)

func TestNothing(t *testing.T) {}

func BenchmarkGopherwoodAdd(b *testing.B) {
	rand.Seed(boringRandomSeed)
	t := Tree{}
	for n := 0; n < b.N; n++ {
		t.Add(randStringBytesRmndr(b.N))
	}
}
