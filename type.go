package gosimdtest

type Float32x4 [4]float32

func AddNoSimd(a, b, r *Float32x4) {
	(*r)[0] = (*a)[0] + (*b)[0]
	(*r)[1] = (*a)[1] + (*b)[1]
	(*r)[2] = (*a)[2] + (*b)[2]
	(*r)[3] = (*a)[3] + (*b)[3]
}
