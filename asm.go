//go:build ignore

package main

import (
	"github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/operand"
)

func main() {
	build.Package("o0olele.github.com/gosimdtest")
	build.TEXT("Add", build.NOSPLIT, "func(x, y, r *Float32x4)")
	var x = build.Load(build.Param("x"), build.GP64())
	var y = build.Load(build.Param("y"), build.GP64())
	var r = build.Load(build.Param("r"), build.GP64())

	var X0 = build.XMM()
	build.MOVUPS(operand.Mem{Base: x}, X0)
	var X1 = build.XMM()
	build.MOVUPS(operand.Mem{Base: y}, X1)
	build.ADDPS(X0, X1)

	build.MOVAPS(X1, operand.Mem{Base: r})

	build.RET()
	build.Generate()
}
