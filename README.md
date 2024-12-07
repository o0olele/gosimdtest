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
# What's Next?
- try to add complex simd funcs like matrix operations and find the differences.
- try on arm cpus.
- try in c/c++.