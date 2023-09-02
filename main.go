package main

import (
	"fmt"
	. "tet3rd/tris"
)

func main() {
	v, c := Uint27(123).Add(MaxUint27-Uint27(55), 0)
	fmt.Println(v, c)
}
