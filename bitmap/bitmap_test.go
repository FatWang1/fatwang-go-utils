package bitmap

import (
	"math"
	"runtime"
	"sync"
	"testing"
)

var (
	bm = NewBitMap(math.MaxInt16 >> 1)
)

func BenchmarkBitMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bm.Set(uint(i % (math.MaxInt16 >> 1)))
	}
}

func BenchmarkBitMap_Check(b *testing.B) {
	bm.Set(1024)
	for i := 0; i < b.N; i++ {
		bm.Check(1024)
	}
}

func BenchmarkBitMap_Del(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bm.Del(uint(i % (math.MaxInt16 >> 1)))
	}
}

func TestQuestion(t *testing.T) {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			println(n)
		}(i)
	}
	wg.Wait()
}
