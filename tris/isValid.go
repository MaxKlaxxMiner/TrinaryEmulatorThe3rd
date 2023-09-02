package tris

func (u Uint9) IsValid() bool {
	return u < Uint9End
}

func (u Uint27) IsValid() bool {
	return u < Uint27End
}

func (s Int9) IsValid() bool {
	return Uint9(s) < Uint9End
}

func (s Int27) IsValid() bool {
	return Uint27(s) < Uint27End
}
