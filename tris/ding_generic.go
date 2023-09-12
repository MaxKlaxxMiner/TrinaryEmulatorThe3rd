//go:build !amd64

package tris

func Ding(val1, val2 uint64) (uint64, uint64) {
	return bits.Mul64(val1, val2)
}

func Gp3(u Uint27) (hi Uint9, mid Uint9, low Uint9) {
	tmp := Uint9(u)
	h, _ := bits.Mul64(uint64(u), 399417452881983571)
	hi = Uint9(h >> 23)
	tmp -= hi * (Uint9End * Uint9End)
	mid = tmp * 3575102585 >> 46
	low = tmp - mid*Uint9End
	return
}
