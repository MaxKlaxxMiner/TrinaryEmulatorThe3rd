package main

import (
	"fmt"
	"math/bits"
	"tet3rd/mem"
	"tet3rd/mem2"
	. "tet3rd/tris"
	"time"
)

// flags: -asmflags='all=-compiling-runtime'

func memtestUint9() {
	m := mem.New(27*27*27 - 10002)
	for i := Uint27(0); i < m.Size; i++ {
		m.SetUint9(i, Uint9(i+10000))
	}

	//fmt.Println(m.Data)

	for i := Uint27(0); i < m.Size; i += 3 {
		fmt.Println(i, "-", m.GetUint9(i+0), m.GetUint9(i+1), m.GetUint9(i+2), "=", i+10000, i+10001, i+10002)
		if m.GetUint9(i+0) != Uint9(i+10000) || m.GetUint9(i+1) != Uint9(i+10001) || m.GetUint9(i+2) != Uint9(i+10002) {
			fmt.Println("ERR")
			break
		}
	}
}

func memtestGetUint9() {
	m := mem.New(27*27*27 - 10002)
	for i := Uint27(0); i < m.Size; i++ {
		m.SetUint9(i, Uint9(i+10000))
	}

	const count = 5000

	for r := 0; r < 3; r++ {
		var sum Uint9
		tim := time.Now()
		for c := 0; c < count; c++ {
			for i := Uint27(0); i < m.Size; i++ {
				sum += m.VanillaGetUint9(i)
			}
		}
		dur := time.Since(tim)
		fmt.Println(Uint27(sum), dur, int(count*float64(m.Size)/float64(dur.Microseconds())), "Ms")
	}
	fmt.Println()
	for r := 0; r < 5; r++ {
		var sum Uint9
		tim := time.Now()
		for c := 0; c < count; c++ {
			for i := Uint27(0); i < m.Size; i++ {
				sum += m.GetUint9(i)
			}
		}
		dur := time.Since(tim)
		fmt.Println(Uint27(sum), dur, int(count*float64(m.Size)/float64(dur.Microseconds())), "Ms")
	}
}

func memtestUint27() {
	m := mem.New(27*27*27 - 10002)
	for i := Uint27(0); i < m.Size; i++ {
		m.SetUint9(i, Uint9(i+10000))
	}

	//fmt.Println(m.Data)

	for i := Uint27(0); i < m.Size-2; i++ {
		fmt.Println(i, "-", m.GetUint27(i), "=", m.VanillaGetUint27(i))
		if m.GetUint27(i) != m.VanillaGetUint27(i) {
			fmt.Println("ERR")
			break
		}
	}
}

func memtestGetUint27() {
	m := mem.New(27*27*27 - 10002)
	for i := Uint27(0); i < m.Size; i++ {
		m.SetUint9(i, Uint9(i+10000))
	}

	const count = 3000

	for r := 0; r < 3; r++ {
		var sum Uint27
		tim := time.Now()
		for c := 0; c < count; c++ {
			for i := Uint27(0); i < m.Size-2; i++ {
				sum += m.VanillaGetUint27(i)
			}
		}
		dur := time.Since(tim)
		fmt.Println(uint64(sum), dur, int(count*float64(m.Size)/float64(dur.Microseconds())), "Ms")
	}
	fmt.Println()
	for r := 0; r < 5; r++ {
		var sum Uint27
		tim := time.Now()
		for c := 0; c < count; c++ {
			for i := Uint27(0); i < m.Size-2; i++ {
				sum += m.GetUint27(i)
			}
		}
		dur := time.Since(tim)
		fmt.Println(uint64(sum), dur, int(count*float64(m.Size)/float64(dur.Microseconds())), "Ms")
	}
}

func memtestGetUint27align() {
	m := mem.New(27*27*27 - 10002)
	for i := Uint27(0); i < m.Size; i++ {
		m.SetUint9(i, Uint9(i+10000))
	}

	const count = 3000

	for r := 0; r < 3; r++ {
		var sum Uint27
		tim := time.Now()
		for c := 0; c < count; c++ {
			for i := Uint27(0); i < m.Size; i += 3 {
				sum += m.VanillaGetUint27(i)
				sum += m.VanillaGetUint27(i)
				sum += m.VanillaGetUint27(i)
			}
		}
		dur := time.Since(tim)
		fmt.Println(uint64(sum), dur, int(count*float64(m.Size)/float64(dur.Microseconds())), "Ms")
	}
	fmt.Println()
	for r := 0; r < 5; r++ {
		var sum Uint27
		tim := time.Now()
		for c := 0; c < count; c++ {
			for i := Uint27(0); i < m.Size; i += 3 {
				sum += m.GetUint27(i)
				sum += m.GetUint27(i)
				sum += m.GetUint27(i)
			}
		}
		dur := time.Since(tim)
		fmt.Println(uint64(sum), dur, int(count*float64(m.Size)/float64(dur.Microseconds())), "Ms")
	}
}

func memtestUint243() {
	m := mem.New(27*27*27 - 10002)
	for i := Uint27(0); i < m.Size; i++ {
		m.SetUint9(i, Uint9(i+10000))
	}

	//fmt.Println(m.Data)

	for i := Uint27(0); i < m.Size-26; i++ {
		var v1, v2 Uint243
		m.GetUint243(i, &v1)
		m.VanillaGetUint243(i, &v2)
		fmt.Println(i, "-", v1, "=", v2)
		if v1 != v2 {
			fmt.Println("ERR")
			break
		}
	}
}

func memtestGetUint243() {
	m := mem2.New(27*27*27 - 10002)
	for i := Uint27(0); i < m.Size; i++ {
		m.SetUint9(i, Uint9(i+10000))
	}

	const count = 3000

	for r := 0; r < 5; r++ {
		var sum Uint27
		tim := time.Now()
		for c := 0; c < count/9; c++ {
			var val Uint243
			for i := Uint27(0); i < m.Size-26; i++ {
				m.GetUint243(i, &val)
				sum += val[0]
				sum += val[1]
				sum += val[2]
				sum += val[3]
				sum += val[4]
				sum += val[5]
				sum += val[6]
				sum += val[7]
				sum += val[8]
			}
		}
		dur := time.Since(tim)
		fmt.Println(uint64(sum), dur, int(count*float64(m.Size)/float64(dur.Microseconds())), "Ms")
	}
}

func memtestGetUint243align() {
	m := mem2.New(27*27*27 - 10002)
	for i := Uint27(0); i < m.Size; i++ {
		m.SetUint9(i, Uint9(i+10000))
	}

	const count = 3000

	for r := 0; r < 5; r++ {
		var sum Uint27
		tim := time.Now()
		for c := 0; c < count/9; c++ {
			var val Uint243
			for i := Uint27(0); i < m.Size-26; i += 3 {
				m.GetUint243(i, &val)
				m.GetUint243(i, &val)
				m.GetUint243(i, &val)
				sum += val[0]
				sum += val[1]
				sum += val[2]
				sum += val[3]
				sum += val[4]
				sum += val[5]
				sum += val[6]
				sum += val[7]
				sum += val[8]
			}
		}
		dur := time.Since(tim)
		fmt.Println(uint64(sum), dur, int(count*float64(m.Size)/float64(dur.Microseconds())), "Ms")
	}
}

func Gp1(u Uint27) (hi Uint9, mid Uint9, low Uint9) {
	tmp := Uint9(u)
	hi = tmp / (Uint9End * Uint9End)
	tmp -= hi * (Uint9End * Uint9End)
	mid = tmp / Uint9End
	low = tmp % Uint9End
	return
}

func Gp2(u Uint27) (hi Uint9, mid Uint9, low Uint9) {
	tmp := Uint9(u)
	hi = tmp / (Uint9End * Uint9End)
	tmp -= hi * (Uint9End * Uint9End)
	mid = tmp * 3575102585 >> 46
	low = tmp - mid*Uint9End
	return
}

func splitTest() {
	const split1 = Uint9End
	const split2 = Uint9End * Uint9End

	//for s3 := Uint9(0); s3 < split1; s3++ {
	//	fmt.Println(s3, "/", int(split1))
	//	for s2 := Uint9(0); s2 < split1; s2++ {
	//		v2 := Uint27(s3*split2 + s2*split1)
	//		for s1 := Uint9(0); s1 < split1; s1++ {
	//			v := v2 + Uint27(s1)
	//			hi, mid, low := v.GetParts()
	//			if hi != s3 || mid != s2 || low != s1 {
	//				fmt.Println("val", s3, s2, s1)
	//				fmt.Println("err", uint64(hi), uint64(mid), uint64(low))
	//				return
	//			}
	//		}
	//	}
	//}

	for s3 := Uint9(0); s3 < split1; s3 += 1970 {
		fmt.Print(s3, " / ", int(split1))
		tim := time.Now()
		sum := Uint9(0)
		for s2 := Uint9(0); s2 < split1; s2++ {
			v2 := Uint27(s3*split2 + s2*split1)
			for s1 := Uint9(0); s1 < split1; s1++ {
				v := v2 + Uint27(s1)
				hi, mid, low := v.GetParts()
				//hi, mid, low := Gp1(v)
				//hi, mid, low := Gp2(v)
				//hi, mid, low := Gp3(v)
				sum += hi + mid + low
				//if hi != s3 || mid != s2 || low != s1 {
				//	fmt.Println()
				//	fmt.Println("val", s3, s2, s1)
				//	fmt.Println("err", uint64(hi), uint64(mid), uint64(low))
				//	break
				//}
			}
		}
		dur := time.Since(tim)
		fmt.Println(" ->", dur, "-", int(sum))
	}

	fmt.Println("ready.")
}

func Ding2(val1, val2 uint64) (uint64, uint64) {
	return bits.Mul64(val1, val2)
}

func main() {
	//memtestUint9()
	//memtestGetUint9()
	//memtestUint27()
	//memtestGetUint27()
	//memtestGetUint27align()
	//memtestUint243()
	//memtestGetUint243()
	//memtestGetUint243align()
	splitTest()
}
