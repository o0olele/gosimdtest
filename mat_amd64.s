//go:build amd64
// +build amd64

#include "textflag.h"

// func Mat4Mul_AVX2_F32(x []float32, y []float32, z []float32)
// Requires: AVX, AVX2, FMA3
TEXT ·Mat4Mul_AVX2_F32(SB), NOSPLIT, $0-72
	MOVQ           x_base+0(FP), DI
	MOVQ           y_base+24(FP), SI
	MOVQ           z_base+48(FP), DX
	VBROADCASTF128 (DX), Y0
	VBROADCASTF128 16(DX), Y1
	VBROADCASTF128 32(DX), Y2
	VBROADCASTF128 48(DX), Y3
	VMOVSS         16(SI), X4
	VMOVSS         (SI), X5
	VSHUFPS        $0x00, X4, X5, X4
	VMOVSS         4(SI), X5
	VMOVSS         8(SI), X6
	VMOVSS         12(SI), X7
	VPERMPD        $0x50, Y4, Y4
	VMULPS         Y4, Y0, Y0
	VMOVSS         20(SI), X4
	VSHUFPS        $0x00, X4, X5, X4
	VPERMPD        $0x50, Y4, Y4
	VFMADD213PS    Y0, Y1, Y4
	VMOVSS         24(SI), X0
	VSHUFPS        $0x00, X0, X6, X0
	VPERMPD        $0x50, Y0, Y0
	VFMADD213PS    Y4, Y2, Y0
	VMOVSS         28(SI), X1
	VSHUFPS        $0x00, X1, X7, X1
	VPERMPD        $0x50, Y1, Y1
	VFMADD213PS    Y0, Y3, Y1
	VBROADCASTF128 (DX), Y0
	VBROADCASTF128 16(DX), Y2
	VBROADCASTF128 32(DX), Y3
	VMOVUPS        Y1, (DI)
	VBROADCASTF128 48(DX), Y1
	VMOVSS         48(SI), X4
	VMOVSS         32(SI), X5
	VSHUFPS        $0x00, X4, X5, X4
	VMOVSS         36(SI), X5
	VMOVSS         40(SI), X6
	VMOVSS         44(SI), X7
	VPERMPD        $0x50, Y4, Y4
	VMULPS         Y4, Y0, Y0
	VMOVSS         52(SI), X4
	VSHUFPS        $0x00, X4, X5, X4
	VPERMPD        $0x50, Y4, Y4
	VFMADD213PS    Y0, Y2, Y4
	VMOVSS         56(SI), X0
	VSHUFPS        $0x00, X0, X6, X0
	VPERMPD        $0x50, Y0, Y0
	VFMADD213PS    Y4, Y3, Y0
	VMOVSS         60(SI), X2
	VSHUFPS        $0x00, X2, X7, X2
	VPERMPD        $0x50, Y2, Y2
	VFMADD213PS    Y0, Y1, Y2
	VMOVUPS        Y2, 32(DI)
	VZEROUPPER
	RET
