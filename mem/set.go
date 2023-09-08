package mem

import (
	"encoding/binary"
	. "tet3rd/tris"
)

func (m *Mem) SetUint9(ofs Uint27, val Uint9) {
	u27ofs := ofs / 3
	num := u27ofs % 8
	bitShift := num * 3
	dataOfs := (u27ofs/8)*43 + num*5
	rawVal := binary.LittleEndian.Uint64(m.Data[dataOfs:])
	u27val := Uint27(rawVal >> bitShift & BitMask)
	rawVal &= ^(BitMask << bitShift)
	if u27val > MaxUint27 {
		panic("value out of range")
	}
	switch ofs % 3 {
	default:
		u27val = val.MergeParts(u27val.GetMid(), u27val.GetHi())
	case 1:
		u27val = u27val.GetLow().MergeParts(val, u27val.GetHi())
	case 2:
		u27val = u27val.GetLow().MergeParts(u27val.GetMid(), val)
	}
	if u27val > MaxUint27 {
		panic("value out of range")
	}
	rawVal |= uint64(u27val) << bitShift
	binary.LittleEndian.PutUint64(m.Data[dataOfs:], rawVal)
}

func (m *Mem) SetUint27(ofs Uint27, val Uint27) {
	m.SetUint9(ofs, val.GetLow())
	m.SetUint9(ofs+1, val.GetMid())
	m.SetUint9(ofs+2, val.GetHi())
}
