package mem

import (
	"encoding/binary"
	. "tet3rd/tris"
)

func (m *Mem) VanillaGetUint9(ofs Uint27) Uint9 {
	if ofs >= m.Size {
		panic("index out of range")
	}
	u27ofs := ofs / 3
	num := u27ofs % 8
	bitShift := num * 3
	dataOfs := (u27ofs/8)*43 + num*5
	u27val := Uint27(binary.LittleEndian.Uint64(m.Data[dataOfs:]) >> bitShift & BitMask)
	if u27val > MaxUint27 {
		panic("value out of range")
	}
	switch ofs % 3 {
	default:
		return u27val.GetLow()
	case 1:
		return u27val.GetMid()
	case 2:
		return u27val.GetHi()
	}
}

func (m *Mem) VanillaGetUint27(ofs Uint27) Uint27 {
	return m.VanillaGetUint9(ofs).MergeParts(m.VanillaGetUint9(ofs+1), m.VanillaGetUint9(ofs+2))
}

func (m *Mem) VanillaGetUint243(ofs Uint27, output *Uint243) {
	for i := range output {
		output[i] = m.VanillaGetUint27(ofs + Uint27(i*3))
	}
}
