package main

import "fmt"

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)
	n = false
	fmt.Printf("%v, %T\n", n, n)

	a := 1 == 1
	b := 1 == 2
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)

	var c bool
	fmt.Printf("%v, %T\n", c, c)

	// int8, int16, int32, int64
	var d int64 = -42
	fmt.Printf("%v, %T\n", d, d)

	// uint8, uint16, uint32, uint64
	var e uint64 = 42
	fmt.Printf("%v, %T\n", e, e)

	x := 10
	y := 3
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	var x1 int = 8
	var x2 int8 = 10
	// fmt.Println(x1 + x2) -- will not work
	fmt.Println(x1 + int(x2))

	x3 := 10              // 1010
	x4 := 3               // 0011
	fmt.Println(x3 & x4)  // 0010 - 2
	fmt.Println(x3 | x4)  // 1011 - 11
	fmt.Println(x3 ^ x4)  // 1001 - 9
	fmt.Println(x3 &^ x4) // 0100 - 8

	x5 := 8              // 2 ^ 3
	fmt.Println(x5 << 3) // 2 ^ 6 = 64
	fmt.Println(x5 >> 3) // 2 ^ 0 = 1

	// float32, float64
	f := 3.14
	f = 13.7e72
	f = 2.1e14
	fmt.Printf("%v, %T\n", f, f)

	// complex64, complex128
	var comp complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", comp, comp)
	fmt.Printf("%v, %T\n", real(comp), real(comp))
	fmt.Printf("%v, %T\n", imag(comp), imag(comp))

	var comp1 complex64 = complex(5, 12)
	fmt.Printf("%v, %T\n", comp1, comp1)

	s := "This is a string"
	s1 := "This is also a string🌎"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2]))
	fmt.Printf("%v, %T\n", s+s1, s+s1)

	// rune = utf32
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
