package tris

func (u Uint18) GetParts() (hi Uint6, mid Uint6, low Uint6) {
	tmp := Uint6(u)
	hi = tmp / (Uint6End * Uint6End)
	tmp -= hi * (Uint6End * Uint6End)
	mid = tmp / Uint6End
	low = tmp % Uint6End
	return
}

func (u Uint18) GetHi() Uint6 {
	return Uint6(u) / (Uint6End * Uint6End)
}

func (u Uint18) GetMid() Uint6 {
	return Uint6(u) / Uint6End % Uint6End
}

func (u Uint18) GetLow() Uint6 {
	return Uint6(u) % Uint6End
}

func (u Uint36) GetParts() (hi Uint18, low Uint18) {
	return u.GetHi(), u.GetLow()
}

func (u Uint36) GetHi() Uint18 {
	return Uint18(u / Uint36(Uint18End))
}

func (u Uint36) GetLow() Uint18 {
	return Uint18(u % Uint36(Uint18End))
}

func (u Uint6) MergeParts(hi, mid Uint6) Uint18 {
	return Uint18(hi)*Uint18(Uint6End*Uint6End) + Uint18(mid)*Uint18(Uint6End) + Uint18(u)
}

func (u Uint18) MergeParts(hi Uint18) Uint36 {
	return Uint36(hi)*Uint36(Uint18End) + Uint36(u)
}
