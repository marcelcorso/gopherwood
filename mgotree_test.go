package gopherwood

import (
	"math/rand"
	"testing"

	"github.com/marcelcorso/gotree"
)

func BenchmarkMGotreeAdd(b *testing.B) {
	rand.Seed(boringRandomSeed)
	t := gotree.MTree{}
	for n := 0; n < b.N; n++ {
		t.Add(randStringBytesRmndr(b.N))
	}
}
