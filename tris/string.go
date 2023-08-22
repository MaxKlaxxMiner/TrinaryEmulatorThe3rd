package tris

import "strconv"

func (u Uint6) String() string {
	if !u.IsValid() {
		return "n/a: " + strconv.FormatUint(uint64(u), 10)
	}
	return strconv.FormatUint(uint64(u), 10)
}

func (u Uint18) String() string {
	if !u.IsValid() {
		return "n/a: " + strconv.FormatUint(uint64(u), 10)
	}
	return strconv.FormatUint(uint64(u), 10)
}

func (u Uint36) String() string {
	if !u.IsValid() {
		return "n/a: " + strconv.FormatUint(uint64(u), 10)
	}
	return strconv.FormatUint(uint64(u), 10)
}

func (s Int6) String() string {
	if !s.IsValid() {
		return "n/a: " + strconv.FormatUint(uint64(s), 10)
	}
	return strconv.FormatInt(s.Int64(), 10)
}

func (s Int18) String() string {
	if !s.IsValid() {
		return "n/a: " + strconv.FormatUint(uint64(s), 10)
	}
	return strconv.FormatInt(s.Int64(), 10)
}

func (s Int36) String() string {
	if !s.IsValid() {
		return "n/a: " + strconv.FormatUint(uint64(s), 10)
	}
	return strconv.FormatInt(s.Int64(), 10)
}
