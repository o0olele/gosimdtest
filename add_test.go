package gosimdtest

import (
	"math/rand/v2"
	"runtime"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	var a = Float32x4{1, 2, 3, 4}
	var b = Float32x4{5, 6.999, 7, 8}
	var c Float32x4
	Add(&a, &b, &c)
	t.Log(c)
}

func BenchmarkADD(b *testing.B) {

	const num = 10000
	var as, bs, rs [num]Float32x4

	for i := 0; i < num; i++ {
		as[i][0], as[i][1], as[i][2], as[i][3] = rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()
		bs[i][0], bs[i][1], bs[i][2], bs[i][3] = rand.Float32(), rand.Float32(), rand.Float32(), rand.Float32()
		rs[i] = Float32x4{}
	}

	b.ResetTimer()
	b.Run("SIMD", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var a = &as[i%num]
			var b = &bs[i%num]
			var c = &rs[i%num]

			Add(a, b, c)
		}
	})

	if !strings.Contains(runtime.GOARCH, "arm") {
		b.Run("SIMD-F4", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rs[i%num] = AddFloat4(as[i%num], bs[i%num])
			}
		})
	}

	b.Run("No-SIMD", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var a = &as[i%num]
			var b = &bs[i%num]
			var c = &rs[i%num]

			AddNoSimd(a, b, c)
		}
	})

	var total float32
	for i := 0; i < num; i++ {
		total += rs[i][0] + as[i][0] + bs[i][0]
	}

	runtime.KeepAlive(&rs)
}
