//go:build arm64
// +build arm64

#include "textflag.h"

// func Add(x *Float32x4, y *Float32x4, r *Float32x4)
TEXT ·Add(SB), NOSPLIT, $0-24
	MOVD	x+0(FP), R0
	MOVD	y+8(FP), R1
	MOVD	r+16(FP), R2
	VLD1.P	16(R0), [V0.S4]
	VLD1.P	16(R1), [V1.S4]
	WORD	$0x4E21D402
	VST1.P	[V2.S4], 16(R2)
	RET
    