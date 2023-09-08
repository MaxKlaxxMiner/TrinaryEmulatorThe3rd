package mem2

import (
	. "tet3rd/tris"
)

func (m *Mem) GetUint9(ofs Uint27) Uint9 {
	return Uint9(m.Data[ofs])
}

func (m *Mem) GetUint27(ofs Uint27) Uint27 {
	d := m.Data
	_ = d[ofs+2]

	return Uint27(Uint9(d[ofs]) + Uint9(d[ofs+1])*Uint9End + Uint9(d[ofs+2])*(Uint9End*Uint9End))
}

func (m *Mem) GetUint243(ofs Uint27, output *Uint243) {
	d := m.Data
	_ = d[ofs+26]

	output[0] = Uint27(Uint9(d[ofs]) + Uint9(d[ofs+1])*Uint9End + Uint9(d[ofs+2])*(Uint9End*Uint9End))
	output[1] = Uint27(Uint9(d[ofs+3]) + Uint9(d[ofs+4])*Uint9End + Uint9(d[ofs+5])*(Uint9End*Uint9End))
	output[2] = Uint27(Uint9(d[ofs+6]) + Uint9(d[ofs+7])*Uint9End + Uint9(d[ofs+8])*(Uint9End*Uint9End))
	output[3] = Uint27(Uint9(d[ofs+9]) + Uint9(d[ofs+10])*Uint9End + Uint9(d[ofs+11])*(Uint9End*Uint9End))
	output[4] = Uint27(Uint9(d[ofs+12]) + Uint9(d[ofs+13])*Uint9End + Uint9(d[ofs+14])*(Uint9End*Uint9End))
	output[5] = Uint27(Uint9(d[ofs+15]) + Uint9(d[ofs+16])*Uint9End + Uint9(d[ofs+17])*(Uint9End*Uint9End))
	output[6] = Uint27(Uint9(d[ofs+18]) + Uint9(d[ofs+19])*Uint9End + Uint9(d[ofs+20])*(Uint9End*Uint9End))
	output[7] = Uint27(Uint9(d[ofs+21]) + Uint9(d[ofs+22])*Uint9End + Uint9(d[ofs+23])*(Uint9End*Uint9End))
	output[8] = Uint27(Uint9(d[ofs+24]) + Uint9(d[ofs+25])*Uint9End + Uint9(d[ofs+26])*(Uint9End*Uint9End))
}
