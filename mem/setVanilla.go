package mem

import (
	"encoding/binary"
	. "tet3rd/tris"
)

func (m *Mem) VanillaSetUint9(ofs Uint27, val Uint9) {
	if ofs >= m.Size {
		panic("index out of range")
	}
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

func (m *Mem) VanillaSetUint27(ofs Uint27, val Uint27) {
	m.VanillaSetUint9(ofs, val.GetLow())
	m.VanillaSetUint9(ofs+1, val.GetMid())
	m.VanillaSetUint9(ofs+2, val.GetHi())
}

func (m *Mem) VanillaSetUint243(ofs Uint27, val *Uint243) {
	m.VanillaSetUint27(ofs, val[0])
	m.VanillaSetUint27(ofs+3, val[1])
	m.VanillaSetUint27(ofs+6, val[2])
	m.VanillaSetUint27(ofs+9, val[3])
	m.VanillaSetUint27(ofs+12, val[4])
	m.VanillaSetUint27(ofs+15, val[5])
	m.VanillaSetUint27(ofs+18, val[6])
	m.VanillaSetUint27(ofs+21, val[7])
	m.VanillaSetUint27(ofs+24, val[8])
}
