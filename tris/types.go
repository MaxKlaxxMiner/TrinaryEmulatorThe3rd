package tris

type Uint9 uint   // 0 - 19682
type Uint27 uint  // 0 - 7625597484986
type Int9 Uint9   // -9841 - +9841
type Int27 Uint27 // -3812798742493 - +3812798742493

const (
	Uint9End = Uint9(3 * 3 * 3 * 3 * 3 * 3 * 3 * 3 * 3)
	MaxUint9 = Uint9End - 1
	MaxInt9  = Int9(MaxUint9) / 2
	MinInt9  = MaxInt9 + 1

	Uint27End = Uint27(Uint9End * Uint9End * Uint9End)
	MaxUint27 = Uint27End - 1
	MaxInt27  = Int27(MaxUint27) / 2
	MinInt27  = MaxInt27 + 1
)
