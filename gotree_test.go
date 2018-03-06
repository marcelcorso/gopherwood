package gopherwood

import "math/rand"

/*
func BenchmarkGotreeAdd(b *testing.B) {
	rand.Seed(boringRandomSeed)
	t := gotree.Tree{}
	for n := 0; n < b.N; n++ {
		t.Add(randStringBytesRmndr(b.N))
	}
}
*/

// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const boringRandomSeed = 10

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
