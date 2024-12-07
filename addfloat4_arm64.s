//go:build arm64
// +build arm64

#include "textflag.h"

// func AddFloat4(a, b [4]float32) [4]float32
TEXT ·Add(SB), NOSPLIT, $0-24
	VLD1.P	x+0(FP), [V0.S4]
	VLD1.P	y+8(FP), [V1.S4]
	WORD	$0x4E21D402
	VST1.P	[V2.S4], 16(r+16(FP))
	RET