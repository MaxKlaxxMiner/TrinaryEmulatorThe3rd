package main

import (
	"fmt"
	. "tet3rd/tris"
)

func main() {
	for i := int32(-100); i <= 100; i++ {
		fmt.Println(i, Int6New(i))
	}
}
