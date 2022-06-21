# go-utils

## Benchmark 说明
```
goos: windows
goarch: amd64
pkg: github.com/alphasnow/go-utils/encode
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkMd5
BenchmarkMd5-12          7277185               161.8 ns/op
PASS
```
- BenchmarkMd5-12: 12核的CPU
- 7277185: 一秒内的函数执行次数
- 161.8 ns/op: 函数每次执行的耗时