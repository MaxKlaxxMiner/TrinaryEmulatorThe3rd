package mem2

import (
	. "tet3rd/tris"
)

type Mem struct {
	Size Uint27   // Größe des reservierten Speichers in Tri-Bytes (Uint9)
	Data []uint16 // reservierter Speicher als uint16
}

func New(size Uint27) *Mem {
	if size < 9 {
		panic("size error (size < 9)")
	}
	if size%3 != 0 {
		panic("size align error (size % 3 != 0)")
	}
	return &Mem{
		Size: size,
		Data: make([]uint16, size), // Speicherplatz reservieren
	}
}
