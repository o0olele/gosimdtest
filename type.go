package gosimdtest

type Float32x4 [4]float32

func AddNoSimd(a, b, r *Float32x4) {
	(*r)[0] = (*a)[0] + (*b)[0]
	(*r)[1] = (*a)[1] + (*b)[1]
	(*r)[2] = (*a)[2] + (*b)[2]
	(*r)[3] = (*a)[3] + (*b)[3]
}

func Mat4MulNoSimd(a, b, dst []float32) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			dst[i*4+j] = a[i*4]*b[j] + a[i*4+1]*b[1*4+j] +
				a[i*4+2]*b[2*4+j] + a[i*4+3]*b[3*4+j]
		}
	}
}
