package prime_seive

import (
	"testing"
)

func TestBrutePrimality(t *testing.T) {
	t.Run("Prime", func(t *testing.T) {
		want := true
		got := BrutePrimality(5)
		assertPrime(t, got, want)
	})
	t.Run("Non-Prime", func(t *testing.T) {
		want := false
		got := BrutePrimality(4)
		assertPrime(t, got, want)
	})
}

func TestSeivePrimality(t *testing.T) {
	t.Run("Prime", func(t *testing.T) {
		want := true
		got := SeivePrimality(5)
		assertPrime(t, got, want)
	})
	t.Run("Non-Prime", func(t *testing.T) {
		want := false
		got := SeivePrimality(4)
		assertPrime(t, got, want)
	})
}

func assertPrime(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf("Got: %v, Want: %v", got, want)
	}
}

func BenchmarkBrutePrimality(b *testing.B) {
	num := 1
	for i := 0; i < b.N; i++ {
		BrutePrimality(num)
		num += 1
	}
}
func BenchmarkSeivePrimality(b *testing.B) {
	num := 1
	for i := 0; i < b.N; i++ {
		SeivePrimality(num)
		num += 1
	}
}

type primeResult struct {
	num     int
	isPrime bool
}

func BenchmarkBrutePrimality2(b *testing.B) {
	primes := make(map[int]bool)
	resultsChannel := make(chan primeResult)
	num := 1
	for i := 0; i < b.N; i++ {
		go func(fakeNum int) {
			curResult := BrutePrimality(fakeNum)
			resultsChannel <- primeResult{fakeNum, curResult}
			r := <-resultsChannel
			primes[r.num] = r.isPrime
		}(num)
		num += 1
	}
}

func BenchmarkSeivePrimality2(b *testing.B) {
	primes := make(map[int]bool)
	resultsChannel := make(chan primeResult)
	num := 1
	for i := 0; i < b.N; i++ {
		go func(fakeNum int) {
			curResult := SeivePrimality(fakeNum)
			resultsChannel <- primeResult{fakeNum, curResult}
			r := <-resultsChannel
			primes[r.num] = r.isPrime
		}(num)
		num += 1
	}
}

/* -------------------------------------------------------------------------
❯ go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/sainad2222/go-learnings/prime-seive
BenchmarkBrutePrimality-10       1000000             33204 ns/op
BenchmarkSeivePrimality-10      24827007               173.0 ns/op
BenchmarkBrutePrimality2-10      1000000              3205 ns/op
BenchmarkSeivePrimality2-10      2205363               692.9 ns/op
PASS
ok      github.com/sainad2222/go-learnings/prime-seive  54.086s

NOT SURE about
1) what 2nd column and -10 at the end of each benchmarck represents
2) why SeivePrimality2 is worse
------------------------------------------------------------------------- */
