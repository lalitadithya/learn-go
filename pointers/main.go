package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 100
	fmt.Println(a, *b)

	a1 := [3]int{1, 2, 3}
	b1 := &a1[0]
	c1 := &a1[1]
	fmt.Printf("%v %p %p\n", a1, b1, c1)

	// no pointer arithmetic

	var ms *myStruct
	ms = new(myStruct)
	(*ms).foo = 293
	ms.foo = 900 // same as above
	fmt.Println(ms)
}
