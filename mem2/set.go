package mem2

import (
	. "tet3rd/tris"
)

func (m *Mem) SetUint9(ofs Uint27, val Uint9) {
	m.Data[ofs] = uint16(val)
}

func (m *Mem) SetUint27(ofs Uint27, val Uint27) {
	m.SetUint9(ofs, val.GetLow())
	m.SetUint9(ofs+1, val.GetMid())
	m.SetUint9(ofs+2, val.GetHi())
}

func (m *Mem) SetUint243(ofs Uint27, val *Uint243) {
	m.SetUint27(ofs, val[0])
	m.SetUint27(ofs+3, val[1])
	m.SetUint27(ofs+6, val[2])
	m.SetUint27(ofs+9, val[3])
	m.SetUint27(ofs+12, val[4])
	m.SetUint27(ofs+15, val[5])
	m.SetUint27(ofs+18, val[6])
	m.SetUint27(ofs+21, val[7])
	m.SetUint27(ofs+24, val[8])
}
