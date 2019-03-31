package main

import (
	"github.com/orcaman/concurrent-map"
	"strconv"
	"sync"
	"testing"
)

type testStruct struct {
	Num int64
	Str string
}

var ts = testStruct{99999, "Let's test concurrent maps performance"}

var result interface{}

type mapWithLock struct {
	M map[string]interface{}
	sync.RWMutex
}

func NewMapWithLock() *mapWithLock {
	return &mapWithLock{M: make(map[string]interface{})}
}

func (m *mapWithLock) Set(key string, value interface{}) {
	m.Lock()
	m.M[key] = value
	m.Unlock()
}

func (m *mapWithLock) Get(key string) (value interface{}, ok bool) {
	m.RLock()
	value, ok = m.M[key]
	m.RUnlock()
	return value, ok
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func BenchmarkLockMapSet(b *testing.B) {
	m := NewMapWithLock()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m.Set(strconv.Itoa(n), ts)
	}
}

func BenchmarkSyncMapSet(b *testing.B) {
	m := sync.Map{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m.Store(strconv.Itoa(n), ts)
	}
}

func BenchmarkConcMapSet(b *testing.B) {
	m := cmap.New()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		m.Set(strconv.Itoa(n), ts)
	}
}

func BenchmarkLockMapGetSame(b *testing.B) {
	m := NewMapWithLock()
	m.Set(ts.Str, ts)
	b.ResetTimer()
	var tsl interface{}
	for n := 0; n < b.N; n++ {
		tsl, _ = m.Get(ts.Str)
	}
	result = tsl
}

func BenchmarkSyncMapGetSame(b *testing.B) {
	m := sync.Map{}
	m.Store(ts.Str, ts)
	b.ResetTimer()
	var tsl interface{}
	for n := 0; n < b.N; n++ {
		tsl, _ = m.Load(ts.Str)
	}
	result = tsl
}

func BenchmarkConcMapGetSame(b *testing.B) {
	m := cmap.New()
	m.Set(ts.Str, ts)
	b.ResetTimer()
	var tsl interface{}
	for n := 0; n < b.N; n++ {
		tsl, _ = m.Get(ts.Str)
	}
	result = tsl
}

func BenchmarkLockMapGet(b *testing.B) {
	m := NewMapWithLock()

	for i := 0; i < 100000; i++ {
		m.Set(strconv.Itoa(i), ts)
	}

	b.ResetTimer()
	var tsl interface{}
	for n := 0; n < b.N; n++ {
		tsl, _ = m.Get(strconv.Itoa(n % 10))
	}
	result = tsl
}

func BenchmarkSyncMapGet(b *testing.B) {
	m := sync.Map{}

	for i := 0; i < 100000; i++ {
		m.Store(strconv.Itoa(i), ts)
	}

	b.ResetTimer()
	var tsl interface{}
	for n := 0; n < b.N; n++ {
		tsl, _ = m.Load(strconv.Itoa(n % 10))
	}
	result = tsl
}

func BenchmarkConcMapGet(b *testing.B) {
	m := cmap.New()

	for i := 0; i < 100000; i++ {
		m.Set(strconv.Itoa(i), ts)
	}

	b.ResetTimer()
	var tsl interface{}
	for n := 0; n < b.N; n++ {
		tsl, _ = m.Get(strconv.Itoa(n % 10))
	}
	result = tsl
}

func BenchmarkLockMapSetParallel(b *testing.B) {
	m := NewMapWithLock()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for n := 0; n < b.N; n++ {
				m.Set(strconv.Itoa(n), ts)
			}
		}
	})
}

func BenchmarkSyncMapSetParallel(b *testing.B) {
	m := sync.Map{}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for n := 0; n < b.N; n++ {
				m.Store(strconv.Itoa(n), ts)
			}
		}
	})
}

func BenchmarkConcMapSetParallel(b *testing.B) {
	m := cmap.New()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for n := 0; n < b.N; n++ {
				m.Set(strconv.Itoa(n), ts)
			}
		}
	})
}

func BenchmarkLockMapGetParallel10(b *testing.B) {
	m := NewMapWithLock()

	for i := 0; i < 100000; i++ {
		m.Set(strconv.Itoa(i), ts)
	}

	b.ResetTimer()
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for n := 0; n < b.N; n++ {
				m.Get(strconv.Itoa(n % 10))
			}
		}
	})
}

func BenchmarkSyncMapGetParallel10(b *testing.B) {
	m := sync.Map{}

	for i := 0; i < 100000; i++ {
		m.Store(strconv.Itoa(i), ts)
	}

	b.ResetTimer()
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for n := 0; n < b.N; n++ {
				m.Load(strconv.Itoa(n % 10))
			}
		}
	})
}

func BenchmarkConcMapGetParallel10(b *testing.B) {
	m := cmap.New()

	for i := 0; i < 100000; i++ {
		m.Set(strconv.Itoa(i), ts)
	}

	b.ResetTimer()
	b.SetParallelism(10)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for n := 0; n < b.N; n++ {
				m.Get(strconv.Itoa(n % 10))
			}
		}
	})
}

func BenchmarkLockMapGetParallel100(b *testing.B) {
	m := NewMapWithLock()

	for i := 0; i < 100000; i++ {
		m.Set(strconv.Itoa(i), ts)
	}

	b.ResetTimer()
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for n := 0; n < b.N; n++ {
				m.Get(strconv.Itoa(n % 10))
			}
		}
	})
}

func BenchmarkSyncMapGetParallel100(b *testing.B) {
	m := sync.Map{}

	for i := 0; i < 100000; i++ {
		m.Store(strconv.Itoa(i), ts)
	}

	b.ResetTimer()
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for n := 0; n < b.N; n++ {
				m.Load(strconv.Itoa(n % 10))
			}
		}
	})
}

func BenchmarkConcMapGetParallel100(b *testing.B) {
	m := cmap.New()

	for i := 0; i < 100000; i++ {
		m.Set(strconv.Itoa(i), ts)
	}

	b.ResetTimer()
	b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for n := 0; n < b.N; n++ {
				m.Get(strconv.Itoa(n % 10))
			}
		}
	})
}
