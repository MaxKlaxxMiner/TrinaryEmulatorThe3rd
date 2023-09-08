package mem

import (
	. "tet3rd/tris"
	"unsafe"
)

type Mem struct {
	Size Uint27  // Größe des reservierten Speichers in Tri-Bytes (Uint9)
	Data []byte  // reservierter Speicher in Bytes (immer ein wenig größer, als benötigt)
	Ptr  uintptr // unsafe-Zeiger auf den Speicherbereich
}

const BitMask uint64 = 0x7ffffffffff // (43 Bits)

func New(size Uint27) *Mem {
	if size < 9 {
		panic("size error (size < 9)")
	}
	if size%3 != 0 {
		panic("size align error (size % 3 != 0)")
	}
	data := make([]byte, (size*43/24+256)/64*64) // Speicherplatz reservieren (großzügig aufrunden)
	result := &Mem{
		Size: size,
		Data: data,
		Ptr:  uintptr(unsafe.Pointer(&data[0])),
	}
	// Speicheradresse auf 64-Byte ausrichten (= Cache-Line)
	// ---> hat meist der Compiler bei der Reservierung bereits beachtet
	for result.Ptr%64 != 0 {
		result.Ptr++
		result.Data = result.Data[1:]
	}
	return result
}
