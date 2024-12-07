//go:build amd64
// +build amd64

#include "textflag.h"

// func AddFloat4(a, b [4]float32) [4]float32
// Requires: SSE
TEXT ·AddFloat4(SB), NOSPLIT, $0-48
    MOVUPS    a+0(FP), X0
    MOVUPS    b+16(FP), X1
    ADDPS   X0, X1
    MOVUPS    X1, ret+32(FP)
    RET
