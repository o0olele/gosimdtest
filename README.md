# What's this?
- I tried to optimze my vector operation using simd, but found the benchmark result was not expected faster.
  Then i post my question on [Reddit](https://www.reddit.com/r/golang/comments/1h6lsnf/why_my_float32x4_add_function_using_simd_assemble/) and found something worth sharing.

- Func Add with SIMD and transform data through pointers.
    ~~~golang
    func Add(x *Float32x4, y *Float32x4, r *Float32x4)

    #include "textflag.h"

    // func Add(x *Float32x4, y *Float32x4, r *Float32x4)
    // Requires: SSE
    TEXT ·Add(SB), NOSPLIT, $0-24
        MOVQ   x+0(FP), AX
        MOVQ   y+8(FP), CX
        MOVQ   r+16(FP), DX
        MOVUPS (AX), X0
        MOVUPS (CX), X1
        ADDPS  X0, X1
        MOVAPS X1, (DX)
        RET
    ~~~
- Func Add without SIMD.
    ~~~golang
    // be careful with bounding checks.
    func AddNoSimd(a, b, r *Float32x4) {
        (*r)[0] = (*a)[0] + (*b)[0]
        (*r)[1] = (*a)[1] + (*b)[1]
        (*r)[2] = (*a)[2] + (*b)[2]
        (*r)[3] = (*a)[3] + (*b)[3]
    }
    ~~~

- Func Add with SIMD and transform data with no pointer.
    ~~~golang
    func AddFloat4(a, b [4]float32) [4]float32

    #include "textflag.h"

    // func AddFloat4(a, b [4]float32) [4]float32
    // Requires: SSE
    TEXT ·AddFloat4(SB), NOSPLIT, $0-48
        MOVUPS    a+0(FP), X0
        MOVUPS    b+16(FP), X1
        ADDPS   X0, X1
        MOVUPS    X1, ret+32(FP)
        RET
    ~~~

- Benchmark
    ~~~
    Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkADD$ o0olele.github.com/gosimdtest

    goos: linux
    goarch: amd64
    pkg: o0olele.github.com/gosimdtest
    cpu: 13th Gen Intel(R) Core(TM) i5-13600KF
    === RUN   BenchmarkADD
    BenchmarkADD
    === RUN   BenchmarkADD/SIMD
    BenchmarkADD/SIMD
    BenchmarkADD/SIMD-20            884831626                1.556 ns/op           0 B/op          0 allocs/op
    === RUN   BenchmarkADD/SIMD-F4
    BenchmarkADD/SIMD-F4
    BenchmarkADD/SIMD-F4-20         464738224                2.538 ns/op           0 B/op          0 allocs/op
    === RUN   BenchmarkADD/No-SIMD
    BenchmarkADD/No-SIMD
    BenchmarkADD/No-SIMD-20         1000000000               1.177 ns/op           0 B/op          0 allocs/op
    PASS
    ok      o0olele.github.com/gosimdtest   4.257s
    ~~~
- the result shows that the SIMD version is slower than the non-SIMD version. As mentioned in the post：
    > @hi65435: Yeah I guess SSE code is only truly a blast when it's larger chunks of code.
    
  this is one possible reason.
# Arm
- Now i tried on arm64 (RaspberryPi 4). the result is:
    ~~~shell
    go test -bench="ADD$" -benchtime=1s .
    goos: linux
    goarch: arm64
    pkg: o0olele.github.com/simd
    BenchmarkADD/SIMD-4         	80144370	        14.91 ns/op
    BenchmarkADD/No-SIMD-4      	99366718	        11.29 ns/op
    PASS
    ok  	o0olele.github.com/simd	3.288s
    ~~~

# Matrix
- matrix mul operation (asm code from [vek](https://github.com/viterin/vek)), the result is:
    ~~~shell
    goos: linux
    goarch: amd64
    pkg: o0olele.github.com/gosimdtest
    cpu: 13th Gen Intel(R) Core(TM) i5-13600KF
    === RUN   BenchmarkMatMul
    BenchmarkMatMul
    === RUN   BenchmarkMatMul/SIMD
    BenchmarkMatMul/SIMD
    BenchmarkMatMul/SIMD-20                 269824773                4.737 ns/op           0 B/op          0 allocs/op
    === RUN   BenchmarkMatMul/No-SIMD
    BenchmarkMatMul/No-SIMD
    BenchmarkMatMul/No-SIMD-20              50856106                24.30 ns/op            0 B/op          0 allocs/op
    PASS
    ok      o0olele.github.com/gosimdtest   2.998s
    ~~~
- now the result is expected faster.

# Array with large size
- i also tried with different size of arrays (use [vek](https://github.com/viterin/vek)), the result is:
    ~~~shell
    goos: windows
    goarch: amd64
    pkg: github.com/viterin/vek/internal/functions
    cpu: AMD Ryzen 7 7840HS w/ Radeon 780M Graphics     
    BenchmarkADD/SIMD-4-16         	353418499	         3.427 ns/op	       0 B/op	       0 allocs/op
    BenchmarkADD/NO-SIMD-4-16      	572597487	         2.174 ns/op	       0 B/op	       0 allocs/op

    BenchmarkADD/SIMD-16-16        	159170994	         7.580 ns/op	       0 B/op	       0 allocs/op
    BenchmarkADD/NO-SIMD-16-16     	136259923	         8.581 ns/op	       0 B/op	       0 allocs/op

    BenchmarkADD/SIMD-32-16        	249333396	         4.937 ns/op	       0 B/op	       0 allocs/op
    BenchmarkADD/NO-SIMD-32-16     	100000000	        15.82 ns/op	       0 B/op	       0 allocs/op

    BenchmarkADD/SIMD-64-16        	133381720	         8.695 ns/op	       0 B/op	       0 allocs/op
    BenchmarkADD/NO-SIMD-64-16     	39249933	        33.63 ns/op	       0 B/op	       0 allocs/op
    ~~~
- it seems that the performance of SIMD only is better than No-SIMD when the size of array is larger.
# What's Next?
- try in c/c++.