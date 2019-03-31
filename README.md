# Concurrent Map Benchmark

Simple benchmark test for https://github.com/orcaman/concurrent-map, https://golang.org/pkg/sync/#Map, and regular map with RWMutex 

Ubuntu 19.10 Intel(R) Core(TM) i3-4350 CPU @ 3.60GHz with 24G of Rams result:

goos: linux  
goarch: amd64  

BenchmarkLockMapSet-4                    2000000               756 ns/op  
BenchmarkSyncMapSet-4                    1000000              1254 ns/op  
BenchmarkConcMapSet-4                    2000000               625 ns/op  

BenchmarkLockMapGetSame-4               100000000               19.8 ns/op  
BenchmarkSyncMapGetSame-4               50000000                36.1 ns/op  
BenchmarkConcMapGetSame-4               20000000                68.3 ns/op  

BenchmarkLockMapGet-4                   50000000                27.4 ns/op  
BenchmarkSyncMapGet-4                   50000000                36.9 ns/op  
BenchmarkConcMapGet-4                   50000000                40.7 ns/op  

BenchmarkLockMapSetParallel-4              10000           2459639 ns/op  
BenchmarkSyncMapSetParallel-4              10000           2977836 ns/op  
BenchmarkConcMapSetParallel-4              10000            970672 ns/op  

BenchmarkLockMapGetParallel10-4            10000            467040 ns/op  
BenchmarkSyncMapGetParallel10-4            10000            170646 ns/op  
BenchmarkConcMapGetParallel10-4            10000            205387 ns/op  

BenchmarkLockMapGetParallel100-4           10000            467468 ns/op  
BenchmarkSyncMapGetParallel100-4           10000            172971 ns/op  
BenchmarkConcMapGetParallel100-4           10000            205827 ns/op  