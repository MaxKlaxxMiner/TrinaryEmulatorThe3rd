package main

import (
	"fmt"
	. "tet3rd/tris"
)

func main() {
	v, c := Uint18(123).Add(MaxUint18-Uint18(55), 0)
	fmt.Println(v, c)
}
