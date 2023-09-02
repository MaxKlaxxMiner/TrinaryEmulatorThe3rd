package tris

import "strconv"

func (u Uint9) String() string {
	if !u.IsValid() {
		return "n/a: " + strconv.FormatUint(uint64(u), 10)
	}
	return strconv.FormatUint(uint64(u), 10)
}

func (u Uint27) String() string {
	if !u.IsValid() {
		return "n/a: " + strconv.FormatUint(uint64(u), 10)
	}
	return strconv.FormatUint(uint64(u), 10)
}

func (s Int9) String() string {
	if !s.IsValid() {
		return "n/a: " + strconv.FormatUint(uint64(s), 10)
	}
	return strconv.Itoa(s.Int())
}

func (s Int27) String() string {
	if !s.IsValid() {
		return "n/a: " + strconv.FormatUint(uint64(s), 10)
	}
	return strconv.Itoa(s.Int())
}
