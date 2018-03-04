package main

import (
	"math/cmplx"
	"fmt"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64-1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main()  {
	const f  = "%T(%v)\n"
	fmt.Print(f, ToBe, ToBe)
	fmt.Print(f, MaxInt, MaxInt)
	fmt.Print(f, z, z)
}