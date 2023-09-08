package mem

import (
	. "tet3rd/tris"
	"unsafe"
)

func (m *Mem) GetUint9(ofs Uint27) Uint9 {
	u27ofs := ofs / 3
	num := u27ofs % 8
	bitShift := num * 3
	dataOfs := (u27ofs/8)*43 + num*5
	u27val := Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
	switch ofs % 3 {
	default:
		return Uint9(u27val) % Uint9End
	case 1:
		return Uint9(u27val) / Uint9End % Uint9End
	case 2:
		return Uint9(u27val) / (Uint9End * Uint9End)
	}
}

func (m *Mem) GetUint27(ofs Uint27) Uint27 {
	u27ofs := ofs / 3
	num := u27ofs % 8
	bitShift := num * 3
	dataOfs := (u27ofs/8)*43 + num*5
	u27val := Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
	if u27ofs*3 == ofs {
		return u27val
	}
	var u27val2 Uint27
	if num == 7 {
		u27val2 = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs+8))) & BitMask)
	} else {
		bitShift += 3
		u27val2 = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs+5))) >> bitShift & BitMask)
	}
	if ofs%3 == 1 {
		return Uint27(Uint9(u27val)/Uint9End + Uint9(u27val2)%Uint9End*(Uint9End*Uint9End))
	} else {
		return Uint27(Uint9(u27val)/(Uint9End*Uint9End) + Uint9(u27val2)%(Uint9End*Uint9End)*Uint9End)
	}
}
