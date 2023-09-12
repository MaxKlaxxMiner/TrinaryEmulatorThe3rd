package tris

func (u Uint27) GetParts() (hi Uint9, mid Uint9, low Uint9) {
	tmp := Uint9(u)
	//h, _ := bits.Mul64(uint64(u), 399417452881983571)
	//hi = Uint9(h >> 23)
	hi = tmp / (Uint9End * Uint9End)
	tmp -= hi * (Uint9End * Uint9End)
	mid = tmp * 3575102585 >> 46
	low = tmp - mid*Uint9End
	return
}

func (u Uint27) GetHi() Uint9 {
	return Uint9(u) / (Uint9End * Uint9End)
}

func (u Uint27) GetMid() Uint9 {
	return Uint9(u) / Uint9End % Uint9End
}

func (u Uint27) GetLow() Uint9 {
	return Uint9(u) % Uint9End
}

func (u Uint9) MergeParts(mid, hi Uint9) Uint27 {
	return Uint27(u + mid*Uint9End + hi*Uint9End*Uint9End)
}
