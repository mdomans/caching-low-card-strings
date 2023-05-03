package caching_low_card_strings

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"runtime"
	"testing"
)

var choices = []string{
	"longstring.longstring.longstring.loingstring.longstring.a",
	"longstring.longstring.longstring.loingstring.longstring.b",
	"longstring.longstring.longstring.loingstring.longstring.c",
	"longstring.longstring.longstring.loingstring.longstring.d",
}

func BenchmarkSmallCache_Set(b *testing.B) {
	cache := NewSmallCache()
	runtime.GC()
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	for i := 0; i < b.N; i++ {
		cache.Set(uuid.New().String(), choices[rand.Intn(len(choices))])
	}
	runtime.ReadMemStats(&m2)
	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}

func BenchmarkNaiveCache_Set(b *testing.B) {
	cache := NewNaiveCache()
	runtime.GC()
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	for i := 0; i < b.N; i++ {
		cache.Set(uuid.New().String(), choices[rand.Intn(len(choices))])
	}
	runtime.ReadMemStats(&m2)
	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}

func BenchmarkHaxCache_Set(b *testing.B) {
	cache := NewHaxCache()
	runtime.GC()
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	for i := 0; i < b.N; i++ {
		cache.Set(uuid.New().String(), choices[rand.Intn(len(choices))])
	}
	runtime.ReadMemStats(&m2)
	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}

func BenchmarkSmallHaxCache_Set(b *testing.B) {
	cache := NewSmallHaxCache()
	runtime.GC()
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	for i := 0; i < b.N; i++ {
		cache.Set(uuid.New().String(), choices[rand.Intn(len(choices))])
	}
	runtime.ReadMemStats(&m2)
	fmt.Println("total:", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Println("mallocs:", m2.Mallocs-m1.Mallocs)
}
