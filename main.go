package main

import (
	"fmt"
	"tet3rd/mem"
	. "tet3rd/tris"
	"time"
)

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

func main() {
	//memtestUint9()
	//memtestGetUint9()
	//memtestUint27()
	//memtestGetUint27()
	memtestGetUint27align()
}
