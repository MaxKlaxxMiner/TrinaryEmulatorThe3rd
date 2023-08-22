package tris

func (s Int6) Signed() bool {
	return s > MaxInt6
}

func (s Int18) Signed() bool {
	return s > MaxInt18
}

func (s Int36) Signed() bool {
	return s > MaxInt36
}

func Int6New(s int32) Int6 {
	if int32(s) < 0 {
		return Int6(s + int32(Uint6End))
	}
	return Int6(s)
}

func Int18New(s int32) Int18 {
	if int32(s) < 0 {
		return Int18(s + int32(Uint18End))
	}
	return Int18(s)
}

func Int36New(s int64) Int36 {
	if int32(s) < 0 {
		return Int36(s + int64(Uint36End))
	}
	return Int36(s)
}

func (s Int6) Int32() int32 {
	if s.Signed() {
		return int32(s) - int32(Uint6End)
	}
	return int32(s)
}

func (s Int6) Int64() int64 {
	if s.Signed() {
		return int64(s) - int64(Uint6End)
	}
	return int64(s)
}

func (s Int18) Int32() int32 {
	if s.Signed() {
		return int32(s) - int32(Uint18End)
	}
	return int32(s)
}

func (s Int18) Int64() int64 {
	if s.Signed() {
		return int64(s) - int64(Uint18End)
	}
	return int64(s)
}

func (s Int36) Int64() int64 {
	if s.Signed() {
		return int64(s) - int64(Uint36End)
	}
	return int64(s)
}
