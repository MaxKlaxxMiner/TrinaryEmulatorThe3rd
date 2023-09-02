package tris

func (u Uint27) GetParts() (hi Uint9, mid Uint9, low Uint9) {
	tmp := Uint9(u)
	hi = tmp / (Uint9End * Uint9End)
	tmp -= hi * (Uint9End * Uint9End)
	mid = tmp / Uint9End
	low = tmp % Uint9End
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

func (u Uint9) MergeParts(hi, mid Uint9) Uint27 {
	return Uint27(hi*Uint9End*Uint9End + mid*Uint9End + u)
}
