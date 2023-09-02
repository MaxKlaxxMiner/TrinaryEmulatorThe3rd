package tris

func (s Int9) Signed() bool {
	return s > MaxInt9
}

func (s Int27) Signed() bool {
	return s > MaxInt27
}

func Int9New(s int) Int9 {
	if s < 0 {
		return Int9(s + int(Uint9End))
	}
	return Int9(s)
}

func Int27New(s int) Int27 {
	if s < 0 {
		return Int27(s + int(Uint27End))
	}
	return Int27(s)
}

func (s Int9) Int() int {
	if s.Signed() {
		return int(s) - int(Uint9End)
	}
	return int(s)
}

func (s Int27) Int() int {
	if s.Signed() {
		return int(s) - int(Uint27End)
	}
	return int(s)
}
