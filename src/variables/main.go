package main

import (
	"fmt"
	"strconv"
)

var n int = 50

var (
	actorName     string = "John"
	companionName string = "Jane"
	doctorNumber  int    = 3
	season        int    = 11
)

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var j int = 43
	fmt.Println(j)

	k := 44
	fmt.Println(k)

	var l float32 = 22.7
	fmt.Printf("%v, %T\n", l, l)

	m := 99.
	fmt.Printf("%v, %T\n", m, m)

	fmt.Println(n)

	var n int = 100
	fmt.Println(n)

	var a int = 76
	fmt.Printf("%v, %T\n", a, a)
	var b float32
	b = float32(a)
	fmt.Printf("%v, %T\n", b, b)

	var c string
	c = strconv.Itoa(a)
	fmt.Printf("%v, %T\n", c, c)
}
