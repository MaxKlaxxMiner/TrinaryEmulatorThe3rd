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

func (m *Mem) GetUint243(ofs Uint27, output *Uint243) {
	if ofs%3 == 0 {
		u27ofs := ofs / 3
		num := u27ofs % 8
		bitShift := num * 3
		dataOfs := (u27ofs/8)*43 + num*5
		output[0] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
		if bitShift == 21 {
			bitShift = 0
			dataOfs += 8
			output[1] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) & BitMask)
		} else {
			bitShift += 3
			dataOfs += 5
			output[1] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
		}
		if bitShift == 21 {
			bitShift = 0
			dataOfs += 8
			output[2] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) & BitMask)
		} else {
			bitShift += 3
			dataOfs += 5
			output[2] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
		}
		if bitShift == 21 {
			bitShift = 0
			dataOfs += 8
			output[3] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) & BitMask)
		} else {
			bitShift += 3
			dataOfs += 5
			output[3] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
		}
		if bitShift == 21 {
			bitShift = 0
			dataOfs += 8
			output[4] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) & BitMask)
		} else {
			bitShift += 3
			dataOfs += 5
			output[4] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
		}
		if bitShift == 21 {
			bitShift = 0
			dataOfs += 8
			output[5] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) & BitMask)
		} else {
			bitShift += 3
			dataOfs += 5
			output[5] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
		}
		if bitShift == 21 {
			bitShift = 0
			dataOfs += 8
			output[6] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) & BitMask)
		} else {
			bitShift += 3
			dataOfs += 5
			output[6] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
		}
		if bitShift == 21 {
			bitShift = 0
			dataOfs += 8
			output[7] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) & BitMask)
		} else {
			bitShift += 3
			dataOfs += 5
			output[7] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
		}
		if bitShift == 21 {
			bitShift = 0
			dataOfs += 8
			output[8] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) & BitMask)
		} else {
			bitShift += 3
			dataOfs += 5
			output[8] = Uint27(*(*uint64)(unsafe.Pointer(m.Ptr + uintptr(dataOfs))) >> bitShift & BitMask)
		}
	} else {
		output[0] = m.GetUint27(ofs)
		output[1] = m.GetUint27(ofs + Uint27(3))
		output[2] = m.GetUint27(ofs + Uint27(6))
		output[3] = m.GetUint27(ofs + Uint27(9))
		output[4] = m.GetUint27(ofs + Uint27(12))
		output[5] = m.GetUint27(ofs + Uint27(15))
		output[6] = m.GetUint27(ofs + Uint27(18))
		output[7] = m.GetUint27(ofs + Uint27(21))
		output[8] = m.GetUint27(ofs + Uint27(24))
	}
}
