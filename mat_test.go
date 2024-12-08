package gosimdtest

import (
	"math/rand/v2"
	"runtime"
	"testing"
)

func TestMatMul(t *testing.T) {
	var a = []float32{
		0, 1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11,
		12, 13, 14, 15,
	}
	var b = []float32{
		2, 0, 0, 0,
		0, 2, 0, 0,
		0, 0, 2, 0,
		0, 0, 0, 2,
	}
	var c = [16]float32{}
	Mat4Mul_AVX2_F32(c[:], a, b)
	t.Log(c)
}

func BenchmarkMatMul(b *testing.B) {
	const num = 10000
	var as, bs, rs [num][]float32

	for i := 0; i < num; i++ {
		as[i] = make([]float32, 16)
		bs[i] = make([]float32, 16)
		rs[i] = make([]float32, 16)
		for j := 0; j < 16; j++ {
			as[i][j] = rand.Float32()
			bs[i][j] = rand.Float32()
		}
	}

	// b.ResetTimer()
	b.Run("SIMD", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Mat4Mul_AVX2_F32(rs[i%num], as[i%num], bs[i%num])
		}
	})

	b.Run("No-SIMD", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Mat4MulNoSimd(as[i%num], bs[i%num], rs[i%num])
		}
	})

	var total float32
	for i := 0; i < num; i++ {
		total += rs[i][0] + as[i][0] + bs[i][0]
	}

	runtime.KeepAlive(&rs)
}
