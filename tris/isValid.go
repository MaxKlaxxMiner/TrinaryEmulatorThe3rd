package tris

func (u Uint6) IsValid() bool {
	return u < Uint6End
}

func (u Uint18) IsValid() bool {
	return u < Uint18End
}

func (u Uint36) IsValid() bool {
	return u < Uint36End
}

func (s Int6) IsValid() bool {
	return Uint6(s) < Uint6End
}

func (s Int18) IsValid() bool {
	return Uint18(s) < Uint18End
}

func (s Int36) IsValid() bool {
	return Uint36(s) < Uint36End
}
