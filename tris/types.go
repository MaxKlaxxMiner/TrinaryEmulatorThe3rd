package tris

type Uint6 uint32  // 0 - 728
type Uint18 uint32 // 0 - 387420488
type Uint36 uint64 // 0 - 150094635296999120
type Int6 Uint6    // -364 - +364
type Int18 Uint18  // -193710244 - +193710244
type Int36 Uint36  // -75047317648499560 - +75047317648499560

const (
	Uint6End  = Uint6(3 * 3 * 3 * 3 * 3 * 3)
	Uint18End = Uint18(Uint6End * Uint6End * Uint6End)
	Uint36End = Uint36(uint64(Uint18End) * uint64(Uint18End))
	MaxUint6  = Uint6End - 1
	MaxUint18 = Uint18End - 1
	MaxUint36 = Uint36End - 1
	MaxInt6   = Int6(MaxUint6) / 2
	MaxInt18  = Int18(MaxUint18) / 2
	MaxInt36  = Int36(MaxUint36) / 2
	MinInt6   = MaxInt6 + 1
	MinInt18  = MaxInt18 + 1
	MinInt36  = MaxInt36 + 1
)
